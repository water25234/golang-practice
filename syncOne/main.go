// https://zhuanlan.zhihu.com/p/44360489

package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

func main() {
	once.Do(onces)
	// time.Sleep(5000)
	once.Do(onced)
}

func example2() {
	for i, v := range make([]string, 10) {
		once.Do(onces)
		fmt.Println("count:", v, "---", i)
	}
	for i := 0; i < 10; i++ {

		go func() {
			once.Do(onced)
			fmt.Println("213")
		}()
	}
	time.Sleep(4000)
}

func onces() {
	fmt.Println("onces")
}

func onced() {
	fmt.Println("onced")
}

func example1() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}

	done := make(chan bool)

	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			fmt.Println(i)
			done <- true
		}()
	}

	for k := 0; k < 10; k++ {
		fmt.Println(k)
		<-done
	}
}
