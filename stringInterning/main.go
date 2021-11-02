// https://kkc.github.io/2020/12/14/golang-string-interning/?fbclid=IwAR361HcnTpRtAY4OQqbAKJklC2bNqBFtSlFtiwblUyXnmWf9LOSTDn13vAw
package main

import (
	"fmt"
	"reflect"
	"sync"
	"unsafe"
)

var (
	pool sync.Pool = sync.Pool{
		New: func() interface{} {
			return make(map[string]string)
		},
	}
)

func pointer(s string) uintptr {
	p := unsafe.Pointer(&s)
	h := *(*reflect.StringHeader)(p)
	return h.Data
}

func main() {
	b := []byte("hello")
	s := string(b)
	t := string(b)
	fmt.Println(pointer(s), pointer(t))
}

// String returns s, interned.
func String(s string) string {
	m := pool.Get().(map[string]string)
	c, ok := m[s]
	if ok {
		pool.Put(m)
		return c
	}
	m[s] = s
	pool.Put(m)
	return s
}

// Bytes returns b converted to a string, interned.
func Bytes(b []byte) string {
	m := pool.Get().(map[string]string)
	c, ok := m[string(b)]
	if ok {
		pool.Put(m)
		return c
	}
	s := string(b)
	m[s] = s
	pool.Put(m)
	return s
}
