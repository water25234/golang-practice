// https://iter01.com/458921.html
package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {

	//建立trace檔案
	f, err := os.Create("1.text")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	//啟動trace goroutine
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	//main
	fmt.Println("Hello World")
}
