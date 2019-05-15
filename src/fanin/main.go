package main

import "fmt"

var (
	N = 5
	M = 5
)

func gen(v string, times int) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for i := 0; i < times; i++ {
			ch <- v
		}
	}()
	return ch
}

func fanIn(times int, inputs []<-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for i := 0; i < times; i++ {
			for _, input := range inputs {
				v := <-input
				ch <- v
			}
		}
	}()
	return ch
}

func main() {
	times := M
	inputs := make([]<-chan string, 0, N)
	for i := 0; i < N; i++ {
		threadName := string('A' + i)
		inputs = append(inputs, gen(threadName, times))
	}
	for char := range fanIn(times, inputs) {
		fmt.Println(char)
	}
}
