package main

import "fmt"

type Message struct {
	msg string
	ch  chan bool
}

func T(msg string) chan Message {
	w := make(chan bool)
	c := make(chan Message)
	go func() {
		defer close(c)
		for i := 0; i < 10; i++ {
			c <- Message{msg: msg, ch: w}
			<-w
		}
	}()
	return c
}

func main() {
	A := T("A")
	B := T("B")
	C := T("C")
	for i := 0; i < 10; i++ {
		a := <-A
		fmt.Println(a.msg)
		b := <-B
		fmt.Println(b.msg)
		c := <-C
		fmt.Println(c.msg)
		a.ch <- true
		b.ch <- true
		c.ch <- true
	}
}
