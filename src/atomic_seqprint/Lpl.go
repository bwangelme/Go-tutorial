package main

import (
	"fmt"
	"sync"
)

type Node struct {
	sig  chan int
	str  string
	prev *Node
}

var (
	num       = 3
	lastOrder = 10
	wg        sync.WaitGroup
)

func main() {
	n1 := Node{make(chan int), "A", nil}
	n2 := Node{make(chan int), "B", &n1}
	n3 := Node{make(chan int), "C", &n2}
	n1.prev = &n3

	wg.Add(3)
	go Print(&n1, 1)
	go Print(&n2, 2)
	go Print(&n3, 3)
	wg.Wait()
}

func Print(node *Node, order int) {
	defer wg.Done()
	for i := 0; i < num; i++ {
		if i == 0 && order == 1 {
			fmt.Println(node.str)
			node.sig <- 1
			continue
		}
		<-node.prev.sig
		fmt.Println(node.str)

		if i == num-1 && order == lastOrder {
			break
		}
		node.sig <- 1
	}
}
