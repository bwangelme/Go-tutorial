package main

import (
	"fmt"
	"log"
	"sync"
)

const endNum = 30

type FairLock struct {
	mu        *sync.Mutex
	cond      *sync.Cond
	isHold    bool
	holdCount int
}

func NewFairLock() sync.Locker {
	mu := new(sync.Mutex)
	cond := sync.NewCond(mu)

	return &FairLock{
		holdCount: 0,
		isHold:    false,
		mu:        mu,
		cond:      cond,
	}
}

func (fl *FairLock) Lock() {
	fl.mu.Lock()
	defer fl.mu.Unlock()

	if !fl.isHold {
		fl.holdCount++
		fl.isHold = true
		return
	}

	fl.holdCount++
	for fl.isHold {
		fl.cond.Wait()
	}
	fl.isHold = true
}

func (fl *FairLock) Unlock() {
	fl.mu.Lock()
	defer fl.mu.Unlock()

	if !fl.isHold {
		log.Fatal("unlock of UnLocked mutex")
	}

	if fl.holdCount > 1 {
		fl.cond.Signal()
	}
	fl.isHold = false
	fl.holdCount--
}

var (
	end = make(chan struct{})
	i   int
)

func threadPrint(threadNum int, threadName string, locker sync.Locker) {
	for i < endNum {
		locker.Lock()
		if i >= endNum {
			locker.Unlock()
			continue
		}
		if i%3 != threadNum {
			locker.Unlock()
			continue
		}

		fmt.Printf("%d: %s\n", i, threadName)
		i += 1
		locker.Unlock()
	}
	end <- struct{}{}
}

func main() {
	mu := NewFairLock()
	names := []string{"A", "B", "C"}

	for idx, name := range names {
		go threadPrint(idx, name, mu)
	}

	for _ = range names {
		<-end
	}
}
