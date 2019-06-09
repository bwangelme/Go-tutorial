package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, exit <-chan struct{}) {
LOOP:
	for {
		select {
		default:
			_ = math.Pi * math.Pi
		case <-exit:
			break LOOP
		}
	}

	wg.Done()
}

func main() {
	n := 10
	var wg sync.WaitGroup
	exit := make(chan struct{})

	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(&wg, exit)
	}

	time.Sleep(time.Second)
	close(exit)
	wg.Wait()
	fmt.Println("Exit all")
}
