package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

//FairLock 公平锁
type FairLock struct {
	isLocked bool
	mu       *sync.Mutex
	signals  []chan struct{}
}

func NewFairLock() *FairLock {
	var signals []chan struct{}

	return &FairLock{
		isLocked: false,
		mu:       new(sync.Mutex),
		signals:  signals,
	}
}

func (fl *FairLock) Lock(signal chan struct{}, id int) {
	fl.mu.Lock()
	if !fl.isLocked {
		fl.isLocked = true
		fl.mu.Unlock()
		return
	}

	fl.signals = append(fl.signals, signal)
	fl.mu.Unlock()

	<-signal
}

func (fl *FairLock) Unlock() {
	fl.mu.Lock()
	defer fl.mu.Unlock()

	if !fl.isLocked {
		log.Fatal("unlock of UnLocked mutex")
	}

	if len(fl.signals) == 0 {
		fl.isLocked = false
	} else {
		signal := fl.signals[0]
		fl.signals = fl.signals[1:]
		signal <- struct{}{}
	}
}

func main() {
	n := 5
	lock := NewFairLock()

	running := make(chan struct{})
	end := make(chan struct{})

	for i := 0; i < n; i++ {
		go func(i int) {
			fmt.Printf("Start Goroutine %d\n", i)

			signal := make(chan struct{})
			running <- struct{}{}
			lock.Lock(signal, i)

			fmt.Printf("Started Goroutine %d\n", i)
			end <- struct{}{}
		}(i)
		<-running
	}
	<-end

	time.Sleep(time.Second * 3)

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
	//Started Goroutine 1
	//Started Goroutine 2
	//Started Goroutine 3
	//Started Goroutine 4
}
