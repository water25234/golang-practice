// https://www.thinbug.com/q/51296922
package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {

	requestChan := make(chan chan error)

	// Starting the service goroutine
	go goroutine(requestChan)

	// Send 5 requests and collect errors
	for i := 0; i < 5; i++ {
		fmt.Printf("Request %v\n", i+1)

		// Make channel that will be used to communicate error back to main
		errorChan := make(chan error)

		requestChan <- errorChan

		fmt.Printf("Erro received: %v\n\n", <-errorChan)

		time.Sleep(1 * time.Second)

	}
}

func goroutine(requestChan <-chan chan error) {

	for {
		select {
		case errChan := <-requestChan:
			fmt.Println("Got request from requestChan")
			errChan <- someOperation()
		}
	}
}

// someOperation that will sometimes return nil error and sometimes it will return error
func someOperation() error {

	if rand.Intn(10) > 5 {
		return nil
	}

	return errors.New("error for someOperation")
}
