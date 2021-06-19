package main

import (
	"fmt"
	"sync"
)

func main() {

	pool := &sync.Pool{

		New: func() interface{} {
			return 100
		},
	}

	init := pool.Get()
	fmt.Println(init)

	// 設定參數
	pool.Put(1)

	// 取得結果
	num := pool.Get()
	fmt.Println(num)

	// 在取得一次是取得default 0
	num = pool.Get()
	fmt.Println(num)
}
