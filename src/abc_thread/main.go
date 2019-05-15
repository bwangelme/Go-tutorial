package main

import (
	"fmt"
	"sync"
)

var i int
var mu sync.Mutex
var endSignal = make(chan struct{})

func threadPrint(threadNum int, threadName string) {
	for {
		mu.Lock()
		if i >= 30 {
			mu.Unlock()
			break
		}

		if i%3 == threadNum {
			fmt.Println(threadName)
			i++
		}
		mu.Unlock()
	}

	endSignal <- struct{}{}
}

func main() {
	names := []string{"A", "B", "C"}

	for idx, name := range names {
		go threadPrint(idx, name)
	}

	for _ = range names {
		<-endSignal
	}
}
