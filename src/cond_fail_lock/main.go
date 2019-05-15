package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type FailLock struct {
	mu        *sync.Mutex
	cond      *sync.Cond
	holdCount int
}

func NewFailLock() *FailLock {
	mu := new(sync.Mutex)
	cond := sync.NewCond(mu)

	return &FailLock{
		holdCount: 0,
		mu:        mu,
		cond:      cond,
	}
}

func (fl *FailLock) Lock() {
	fl.mu.Lock()
	defer fl.mu.Unlock()

	fl.holdCount++
	if fl.holdCount == 1 {
		return
	}

	fl.cond.Wait()
}

func (fl *FailLock) Unlock() {
	fl.mu.Lock()
	defer fl.mu.Unlock()

	if fl.holdCount == 0 {
		log.Fatal("unlock of UnLocked mutex")
	}

	fl.holdCount--
	if fl.holdCount != 0 {
		fl.cond.Signal()
	}
}

func main() {
	n := 5
	lock := NewFailLock()

	running := make(chan struct{}) // running 用来保证按顺序启动 N 个 Goroutine
	end := make(chan struct{})     // end 用来保证只有前一个 Goroutine 运行完成后，才会去解锁下一个 Goroutine

	for i := 0; i < n; i++ {
		go func(i int) {
			fmt.Printf("Start Goroutine %d\n", i)
			running <- struct{}{}

			lock.Lock()
			fmt.Printf("Started Goroutine %d\n", i)
			end <- struct{}{}
		}(i)
		<-running
	}
	<-end

	time.Sleep(time.Second * 1)
	fmt.Println("Sleeping 1 second")

	for i := 0; i < n; i++ {
		lock.Unlock()
		if i != n-1 {
			<-end
		}
	}
	//Out
	//Start Goroutine 0
	//Started Goroutine 0
	//Start Goroutine 1
	//Start Goroutine 2
	//Start Goroutine 3
	//Start Goroutine 4
	//Sleeping 1 second
	//Started Goroutine 1
	//Started Goroutine 2
	//Started Goroutine 3
	//Started Goroutine 4
}
