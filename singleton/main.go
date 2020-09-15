// https://www.twblogs.net/a/5db39e0cbd9eee310ee6a098
// https://www.jianshu.com/p/a21d17fc5598
// http://researchlab.github.io/2016/02/25/singleton-pattern-in-go/

package main

import (
	"fmt"
)

type singleton struct {
	count int
}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) Addone() int {
	s.count++
	return s.count
}

func example1() {
	getInstance1 := GetInstance()
	fmt.Println(getInstance1.count)
	addone1 := getInstance1.Addone()
	fmt.Println(addone1)

	getInstance2 := GetInstance()
	fmt.Println(getInstance2.count)
	addone2 := getInstance2.Addone()
	fmt.Println(addone2)

	if addone1 != addone2 {
		fmt.Println("instance is not equal。")
	}
}

func main() {
	example1()
}
