package main

import "fmt"

func main() {
	ch := make(chan int)
	//go func() {
	ch <- 1 // 等到天荒地老
	//}()

	fmt.Println(<-ch)
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
		ch <- 1 // 等到天荒地老
	}()

	fmt.Println(<-ch)
}
