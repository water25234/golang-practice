// https://peterhpchen.github.io/2020/03/08/goroutine-and-channel.html
// https://ithelp.ithome.com.tw/articles/10212068
// https://www.jishuwen.com/d/2SFS/zh-tw
// https://liushuchun.gitbooks.io/golang/content/go_concurence.html

package main

import (
	"fmt"
	"sync"
	"time"
)

// 1.日常在使用 channel 中，要注意區分有緩衝( buffered channel ，
// 	非同步佇列-FIFO處理)與無緩衝( unbuffered channel ，同步流入流出)通道的區別，掌握各自適合使用的方式；
// 2. 出現 deadlock 一定是執行緒/協程之間存在了資源競爭，互相佔用對方需要的資源導致程式永遠不能退出，需要小心可能遇到的坑，
// 	也可以通過加鎖避免。

func main() {
	example1()
}

// channel is deadlock
func example1() {
	ch := make(chan int)
	ch <- 1 // 等到天荒地老
	fmt.Println(<-ch)
}

// solve problem with example1, use goroutine
func example2() {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()

	fmt.Println(<-ch)
}

// channel is deadlock
func example3() {
	ch := make(chan int, 1)
	<-ch
}

// solve problem with example3, use select case
func example4() {
	ch := make(chan int, 1)

	select {
	case v := <-ch:
		fmt.Println(v)
	default:
		fmt.Println("chan on data")
	}
}

// 有緩衝 chan 超過容量時產生死鎖
func example5() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	// ch <- 4 deadlock, ch over then 3
	for v := range ch {
		fmt.Println(v)
	}
}

func example6() {
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3

	// defer close(ch) // 解決方式1：關閉chan

	// range 一直讀取直到chan關閉，否則產生阻塞死鎖

	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()

	// 解決方式2：開啟子協程，主協程sleep等待
	// time.Sleep(1e9)
}

type SafeNumber struct {
	v   int
	mux sync.Mutex
}

func example7() {
	total := SafeNumber{
		v: 0,
	}
	for i := 0; i < 5; i++ {
		go func() {
			total.mux.Lock()
			total.v++
			total.mux.Unlock()
		}()
	}
	time.Sleep(time.Second)
	total.mux.Lock()
	fmt.Println(total.v)
	total.mux.Unlock()
}

func example8() {
	ch := make(chan int, 1)
	ch <- 0
	for i := 0; i < 1000; i++ {
		go func() {
			ch <- <-ch + 1
		}()
		//fmt.Println(1)
	}

	fmt.Println(<-ch)
}

func say(s string, c chan string) {
	for i := 0; i < 5; i++ {
		fmt.Println(i + 1)
	}
	c <- "FINISH"
}

func example9() {
	ch := make(chan string)

	go say("world", ch)
	go say("hello", ch)
	go say("Justin", ch)
	go say("Stacy", ch)
	go say("Gill", ch)

	fmt.Println(len(ch), "123")
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(len(ch), "456")
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(len(ch), "789")
}
