package main

import (
	"fmt"
)

func main() {
	run := func(str string, wait chan struct{}, sig chan struct{}) {
		for i := 0; i < 10; i++ {
			<-wait
			fmt.Printf("%s\n", str)
			sig <- struct{}{}
		}
	}
	c1 := make(chan struct{})
	c2 := make(chan struct{})
	c3 := make(chan struct{})
	c4 := make(chan struct{})
	go run("A", c1, c2)
	go run("B", c2, c3)
	go run("C", c3, c4)
	c1 <- struct{}{}
	for i := 0; i < 10; i++ {
		<-c4
		if i < 9 {
			c1 <- struct{}{}
		}
	}
}
