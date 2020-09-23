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
			singleton.Addone()
			fmt.Println(singleton.count)
		}()
	}
	time.Sleep(time.Second * 2)
}

type singleton struct {
	count int
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

func (s *singleton) Addone() int {
	s.count++
	return s.count
}
