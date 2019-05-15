package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func output(s string, n int, recv, send chan bool) {

	for i := 0; i < n; i++ {
		if nil != recv {
			<-recv
		}
		fmt.Println(i, ":", s)
		if nil != send {
			send <- true
		}

	}
	wg.Done()
}

func coordinator(n int, send, recv chan bool) {

	for i := 0; i < n; i++ {
		send <- true
		<-recv
	}
	wg.Done()
}

func stream() {
	ch := make(chan bool)
	cha := make(chan bool)
	chb := make(chan bool)
	chc := make(chan bool)
	go coordinator(10, ch, chc)
	go output("A", 10, ch, cha)
	go output("B", 10, cha, chb)
	go output("C", 10, chb, chc)
}

func main() {
	wg.Add(4)
	stream()
	wg.Wait()
}
