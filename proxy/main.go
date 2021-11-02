package main

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"hash"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/hauke96/sigolo"
)

const LISTEN = `:8081`
const CACHEDIR = `./cache/`

var cache *Cache

// Hop-by-hop headers. These are removed when sent to the backend.
// http://www.w3.org/Protocols/rfc2616/rfc2616-sec13.html
var hopHeaders = []string{
	"Connection",
	"Keep-Alive",
	"Proxy-Authenticate",
	"Proxy-Authorization",
	"Te", // canonicalized version of "TE"
	"Trailers",
	"Transfer-Encoding",
	"Upgrade",
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

func delHopHeaders(header http.Header) {
	for _, h := range hopHeaders {
		header.Del(h)
	}
}

func appendHostToXForwardHeader(header http.Header, host string) {
	// If we aren't the first proxy retain prior
	// X-Forwarded-For information as a comma+space
	// separated list and fold multiple headers into one.
	if prior, ok := header["X-Forwarded-For"]; ok {
		host = strings.Join(prior, ", ") + ", " + host
	}
	header.Set("X-Forwarded-For", host)
}

type proxy struct {
}

func handleError(err error, w http.ResponseWriter) {
	fmt.Println(err.Error())
	w.WriteHeader(500)
	fmt.Fprintf(w, err.Error())
}

func (p *proxy) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	log.Println(req.RemoteAddr, " ", req.Method, " ", req.URL)
	if req.URL.Scheme == `` {
		if req.URL.Port() == `443` {
			req.URL.Scheme = "https"
			req.URL.Host = req.URL.Hostname()
		} else {
			req.URL.Scheme = "http"
		}
	}

	fullUrl := req.URL.String()
	fmt.Println(fullUrl)

	client := &http.Client{}

	//http: Request.RequestURI can't be set in client requests.
	//http://golang.org/src/pkg/net/http/client.go
	req.RequestURI = ""

	delHopHeaders(req.Header)

	if (req.Method == http.MethodGet || req.Method == http.MethodConnect) && cache.has(fullUrl) {

		// if req.Method == http.MethodConnect {
		// 	req.Method = http.MethodGet
		// }

		fmt.Println(fullUrl)
		content, err := cache.get(fullUrl)
		if err != nil {
			handleError(err, wr)
		} else {
			wr.Write(content)
		}
	}

	if clientIP, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
		appendHostToXForwardHeader(req.Header, clientIP)
	}
	fmt.Println(req)
	resp, err := client.Do(req)
	if err != nil {
		http.Error(wr, "Server Error", http.StatusInternalServerError)
		log.Fatal("ServeHTTP:", err)
	}

	defer resp.Body.Close()
	log.Println(req.RemoteAddr, " ", resp.Status)

	if req.Method == http.MethodGet {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			handleError(err, wr)
			return
		}
		err = cache.put(fullUrl, body)
		if err != nil {
			fmt.Printf("Failed write into cache: %s\n", err)
		}
		wr.Write(body)
	}

	delHopHeaders(resp.Header)

	copyHeader(wr.Header(), resp.Header)
	wr.WriteHeader(resp.StatusCode)
	if req.Method != http.MethodGet {
		io.Copy(wr, resp.Body)
	}

}

type Cache struct {
	folder      string
	hash        hash.Hash
	knownValues map[string][]byte
	mutex       *sync.Mutex
}

func CreateCache(path string) (*Cache, error) {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		sigolo.Error("Cannot open cache folder '%s': %s", path, err)
		sigolo.Info("Create cache folder '%s'", path)
		os.Mkdir(path, os.ModePerm)
	}

	values := make(map[string][]byte, 0)

	// Go through every file an save its name in the map. The content of the file
	// is loaded when needed. This makes sure that we don't have to read
	// the directory content each time the user wants data that's not yet loaded.
	for _, info := range fileInfos {
		if !info.IsDir() {
			values[info.Name()] = nil
		}
	}

	hash := sha256.New()

	mutex := &sync.Mutex{}

	cache := &Cache{
		folder:      path,
		hash:        hash,
		knownValues: values,
		mutex:       mutex,
	}

	return cache, nil
}

func (c *Cache) has(key string) bool {
	hashValue := calcHash(key)

	c.mutex.Lock()
	_, ok := c.knownValues[hashValue]
	c.mutex.Unlock()

	return ok
}

func (c *Cache) get(key string) ([]byte, error) {
	hashValue := calcHash(key)

	// Try to get content. Error if not found.
	c.mutex.Lock()
	content, ok := c.knownValues[hashValue]
	c.mutex.Unlock()
	if !ok {
		sigolo.Debug("Cache doen't know key '%s'", hashValue)
		return nil, errors.New(fmt.Sprintf("Key '%s' is not known to cache", hashValue))
	}

	sigolo.Debug("Cache has key '%s'", hashValue)

	// Key is known, but not loaded into RAM
	if content == nil {
		sigolo.Debug("Cache has content for '%s' already loaded", hashValue)

		content, err := ioutil.ReadFile(c.folder + hashValue)
		if err != nil {
			sigolo.Error("Error reading cached file '%s': %s", hashValue, err)
			return nil, err
		}

		c.mutex.Lock()
		c.knownValues[hashValue] = content
		c.mutex.Unlock()
	}

	return content, nil
}

func (c *Cache) put(key string, content []byte) error {
	hashValue := calcHash(key)

	err := ioutil.WriteFile(c.folder+hashValue, content, 0644)

	// Make sure, that the RAM-cache only holds values we were able to write.
	// This is a decision to prevent a false impression of the cache: If the
	// write fails, the cache isn't working correctly, which should be fixed by
	// the user of this cache.
	if err == nil {
		sigolo.Debug("Cache wrote content into '%s'", hashValue)
		c.mutex.Lock()
		c.knownValues[hashValue] = content
		c.mutex.Unlock()
	}

	return err
}

func calcHash(data string) string {
	sha := sha256.Sum256([]byte(data))
	return hex.EncodeToString(sha[:])
}

func main() {
	addr := flag.String("LISTEN", LISTEN, "The LISTEN of the application.")
	cacheDir := flag.String("CACHEDIR", CACHEDIR, "The CACHE DIRectory, please end with /")
	flag.Parse()
	var err error
	cache, err = CreateCache(*cacheDir)
	if err != nil {
		fmt.Println(`Unable to create cache directory: ` + *cacheDir)
		return
	}

	handler := &proxy{}

	log.Println("Starting proxy server on", *addr)
	if err := http.ListenAndServe(*addr, handler); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
