package main

import (
	"fmt"
	"time"

	"github.com/viney-shih/goroutines"
)

func main() {
	taskN := 100
	rets := make(chan int, taskN)

	// allocate a pool with max size 10.
	// at the beginning, the pool initialized the queue with size 5
	//   and 1 worker to deal with those tasks.
	p := goroutines.NewPool(
		10,
		goroutines.WithTaskQueueLength(5),
		goroutines.WithPreAllocWorkers(1),
	)
	// don't forget to release the pool in the end
	defer p.Release()

	// assign tasks to asynchronous goroutine pool
	for i := 0; i < taskN; i++ {
		idx := i
		p.Schedule(func() {
			// sleep and return the index
			time.Sleep(100 * time.Millisecond)
			rets <- idx
		})
	}

	// wait until all tasks done
	for i := 0; i < taskN; i++ {
		fmt.Println("index:", <-rets)
	}

	// Output: (the order is not the same with input one)
	// index: 3
	// index: 1
	// index: 2
	// index: 4
	// ...
}
