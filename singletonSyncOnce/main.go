package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//啓動多個協程獲取對象並且調用對象的Print內存地址的方法
	for i := 0; i < 5; i++ {
		go func() {
			singleton := GetSingleton()
			singleton.PrintAddress()
		}()
	}
	time.Sleep(time.Second)
}

type singleton struct {
}

var st *singleton
var once sync.Once

func GetSingleton() *singleton {
	if st != nil {
		return st
	}
	once.Do(func() {
		time.Sleep(time.Millisecond * 10)
		st = &singleton{}
	})
	return st
}

func (s *singleton) PrintAddress() {
	fmt.Printf("%p \n", s)
}
