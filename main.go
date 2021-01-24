package main

import (
	"fmt"
	"reflect"
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

func runCalcul() {
	fmt.Println("We have $1500")
	wg.Add(2)
	go withdraw(1) // first withdraw
	go withdraw(2) // second withdraw
	wg.Wait()
}

type Service interface {
	Run() error

	Destroy() error

	Get() interface{}
}

type DB struct {
	ReadDB  string
	WriteDB string
}

func getName(sv Service) string {
	t := reflect.TypeOf(sv)
	if t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	}
	return t.Name()
}

func (db *DB) Run() error {
	return nil
}

func (db *DB) Destroy() error {
	return nil
}

func (db *DB) Get() interface{} {
	return nil
}

type T int

func IsClosed(ch <-chan T) bool {

	select {
	case <-ch:
		return true
	default:
		fmt.Println("default")
	}

	return false
}

const (
	createTimeout = 2000 * time.Millisecond
)

func main() {

	done := make(chan struct{}, 1)
	go func() {
		time.Sleep(10000 * time.Millisecond)
		done <- struct{}{}
		// <-done
	}()
	select {
	case <-done:
		// case done <- struct{}{}:

		fmt.Println("done")
	case <-time.After(4000 * time.Millisecond):
		fmt.Println(fmt.Sprintf("create service failed, time out %s", createTimeout))
	}

	// c := make(chan T)
	// fmt.Println(IsClosed(c)) // false
	// close(c)
	// fmt.Println(IsClosed(c)) // true

	// fmt.Println(getName(&DB{}))

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
