package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/proxy"
)

func main() {
	fmt.Println(1234)
	// Instantiate default collector
	c := colly.NewCollector(colly.AllowURLRevisit())

	// Rotate two socks5 proxies
	// rp, err := proxy.RoundRobinProxySwitcher("socks5://127.0.0.1:1337", "socks5://127.0.0.1:1338")
	// curl -x "http://scraperapi:517b9015124ae652f84194455b192b63@proxy-server.scraperapi.com:8001" -k "http://httpbin.org/ip"
	// curl "http://api.scraperapi.com?api_key=517b9015124ae652f84194455b192b63&url=http://httpbin.org/ip"

	rp, err := proxy.RoundRobinProxySwitcher("http://127.0.0.1:8081", "http://127.0.0.1:8082")
	if err != nil {
		log.Fatal(err)
	}
	// 【设置代理IP】 ，这里使用的是轮询ip方式
	c.SetProxyFunc(rp)

	// Print the response
	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r)
		fmt.Printf("Proxy Address: %s\n", r.Request.ProxyURL)
		fmt.Printf("%s\n", bytes.Replace(r.Body, []byte("\n"), nil, -1))
	})

	// Fetch httpbin.org/ip five times
	for i := 1; i < 6; i++ {
		c.Visit("https://httpbin.org/ip")
		fmt.Println(i)
	}
}
