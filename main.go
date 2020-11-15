package main

import (
	"fmt"
	// "time"
	"sync"
)

type safeNumber struct {
	total int
	mutx  sync.Mutex
}

func main() {

	// fmt.Println(<-ch) // 被上面阻塞，無法被執行到
	// i := 0
	// for i = 0; i < 3; i++ {
	// 	go func(i int) {
	// 		fmt.Println(i)
	// 	}(i)
	// }

	// lock := safeNumber{
	// 	total: 0,
	// }

	// c := make(chan bool, 1)

	// // 以下code 會有什麼問題？如何改善？
	// for i := 0; i < 3; i++ {
	// 	go func() {
	// 		// lock.mutx.Lock()
	// 		fmt.Println(i)
	// 		c <- true
	// 		// lock.mutx.Unlock()
	// 	}()
	// 	<- c
	// 	// time.Sleep(1 * time.Second)
	// }
	// time.Sleep(1 * time.Second)

	// m := make(map[string]int)
	// go func() {
	// 	for {
	// 		v1 := m["k1"]
	// 		fmt.Println("v1: ", v1)
	// 		time.Sleep(1 * time.Second)
	// 	}
	// }()
	// go func() {
	// 	for {
	// 		m["k1"] = 2
	// 		time.Sleep(100 * time.Millisecond)
	// 	}
	// }()
	// for {
	// 	time.Sleep(10 * time.Second)
	// }

	// slice := make([]int)

	// defer func() {
	// 	Done()
	// 	fmt.Println("a")
	// }()
	chA := make(chan func() string, 1)
	// chB := make(chan string, 1)
	chA <- ToDo
	// chB <- ToDo()
	// fmt.Println("b")
	// fmt.Println(<-chB)
	fmt.Println((<-chA)())
}

// func Done() {
// 	fmt.Println("s")
// }
func ToDo() string {
	defer fmt.Println("y")
	return "t"
}

// y
// b
// t
// y
// t
// s
// a
