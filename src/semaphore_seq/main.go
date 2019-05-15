package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/sync/semaphore"
)

var (
	N   = 5
	M   = 3
	ctx = context.TODO()
)

func main() {
	var wait, sig, firstWait, lastSig *semaphore.Weighted

	wait = semaphore.NewWeighted(0)
	firstWait = wait

	for i := 0; i < N; i++ {
		sig = semaphore.NewWeighted(0)
		lastSig = sig
		go echo(i, wait, sig)
		wait = sig
	}

	time.Sleep(time.Second * 1)
	for i := 0; i < M; i++ {
		firstWait.Release(1)
		if err := lastSig.Acquire(ctx, 1); err != nil {
			log.Fatalln("Failed to acquire semaphore: %v", err)
		}
	}

	if err := lastSig.Acquire(ctx, 1); err != nil {
		log.Fatalln("Failed to acquire semaphore: %v", err)
	}

}

func echo(threadNum int, wait *semaphore.Weighted, sig *semaphore.Weighted) {
	threadName := string('A' + threadNum)

	for i := 0; i < M; i++ {
		if err := wait.Acquire(ctx, 1); err != nil {
			log.Fatalln("Failed to acquire semaphore: %v", err)
		}
		fmt.Printf("%d: %s\n", threadNum, threadName)
		sig.Release(1)
	}

	sig.Release(1)
}
