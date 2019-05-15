package main

import (
	"fmt"
	"log"
)

var (
	N = 5
	M = 3
)

func main() {
	var wait, sig, firstWait, lastSig chan struct{}

	wait = make(chan struct{})
	firstWait = wait

	for i := 0; i < N; i++ {
		sig = make(chan struct{})
		lastSig = sig
		go echo(i, wait, sig)
		wait = sig
	}

	for i := 0; i < M; i++ {
		firstWait <- struct{}{}
		<-lastSig
	}
	close(firstWait)

	_, ok := <-lastSig
	if ok {
		log.Fatalln("Channel not closed")
	}
}

func echo(threadNum int, wait chan struct{}, sig chan struct{}) {
	threadName := string('A' + threadNum)

	for _ = range wait {
		fmt.Printf("%d: %s\n", threadNum, threadName)
		sig <- struct{}{}
	}

	fmt.Println("Close", threadName)
	close(sig)
}
