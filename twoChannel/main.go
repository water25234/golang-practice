// http://tleyden.github.io/blog/2013/11/23/understanding-chan-chans-in-go/
package main

import (
	"fmt"
	"time"
)

func main() {

	// make the request chan chan that both go-routines will be given
	requestChan := make(chan chan string)

	// start the goroutines
	go goroutineC(requestChan)
	go goroutineD(requestChan)

	// sleep for a second to let the goroutines complete
	time.Sleep(time.Second)

}

func goroutineC(requestChan chan chan string) {

	// make a new response chan
	responseChan := make(chan string)

	// send the responseChan to goRoutineD
	requestChan <- responseChan

	// read the response
	fmt.Printf("Response: %v\n", <-responseChan)

}

func goroutineD(requestChan chan chan string) {

	// read the responseChan from the requestChan
	responseChan := <-requestChan

	// send a value down the responseChan
	responseChan <- "wassup!"

}
