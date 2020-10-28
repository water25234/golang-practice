package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var money int = 1500
var mu sync.Mutex

func calc(index string, a, realb, b int) int {
	ret := a + b
	// fmt.Println(index, a, b, ret)
	fmt.Println(index, a, realb, b, ret)
	return ret
}

func withdraw(i int) {
	{
		mu.Lock()
		balance := money
		time.Sleep(3000 * time.Millisecond)
		balance -= 1000
		money = balance
		fmt.Println("After withdrawing $1000, balace: ", money, i)
		mu.Unlock()
	}

	wg.Done()
}

func main() {

	fmt.Println("We have $1500")
	wg.Add(2)
	go withdraw(1) // first withdraw
	go withdraw(2) // second withdraw
	wg.Wait()

	// a := 1
	// b := 2
	// defer calc("1", a, b, calc("10", a, b, b))
	// a = 0
	// defer calc("2", a, b, calc("20", a, b, b))
	// b = 1 // not work

	// runtime.GOMAXPROCS(1)
	// wg := sync.WaitGroup{}
	// wg.Add(20)

	// for i := 0; i < 10; i++ {
	// 	go func() {
	// 		fmt.Println("run 1: ", i)
	// 		wg.Done()
	// 	}()
	// }

	// for i := 0; i < 10; i++ {
	// 	go func(i int) {
	// 		fmt.Println("run 2: ", i)
	// 		// wg.Done()
	// 	}(i)
	// }

	// wg.Wait()
}
