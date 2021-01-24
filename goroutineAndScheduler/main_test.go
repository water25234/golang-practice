// go test -v -bench=. -run=none .
// go test -v -bench=. -run=none -benchmem .
// https://blog.wu-boy.com/2018/06/how-to-write-benchmark-in-go/
// https://medium.com/@genchilu/%E7%95%B6%E4%B8%80%E5%80%8B-goroutine-%E5%89%B5%E5%BB%BA%E6%96%B0%E7%9A%84-goroutine-%E6%99%82-scheduler-%E6%9C%83%E9%81%B8%E8%AA%B0-257f434ee1bf
package main

import (
	"testing"
)

func consume(in <-chan interface{}) <-chan interface{} {
	finishe := make(chan interface{})
	go func() {
		defer close(finishe)
		select {
		case <-in:
		}
	}()
	return finishe
}

func produce(out chan<- interface{}) {
	var sum uint64
	for i := uint64(0); i < 100; i++ {
		sum += i * i
	}
	out <- sum
}

var numberOfConsumer int = 100000

func BenchmarkProducerFirst(b *testing.B) {
	buffer := make(chan interface{}, numberOfConsumer)
	finishChans := make([]<-chan interface{}, numberOfConsumer, numberOfConsumer)

	for i := 0; i < numberOfConsumer; i++ {
		produce(buffer)
	}

	for i := 0; i < numberOfConsumer; i++ {
		finishChans[i] = consume(buffer)
	}

	for i := 0; i < numberOfConsumer; i++ {
		<-finishChans[i]
	}
}

func BenchmarkComsumerFirst(b *testing.B) {
	buffer := make(chan interface{}, numberOfConsumer)
	finishChans := make([]<-chan interface{}, numberOfConsumer, numberOfConsumer)

	for i := 0; i < numberOfConsumer; i++ {
		finishChans[i] = consume(buffer)
	}

	for i := 0; i < numberOfConsumer; i++ {
		produce(buffer)
	}

	for i := 0; i < numberOfConsumer; i++ {
		<-finishChans[i]
	}
}
