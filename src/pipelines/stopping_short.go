package main

import "fmt"

func gen(nums ...int) <-chan int {
	out := make(chan int, len(nums))
	for _, n := range nums {
		out <- n
	}
	close(out)
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()

	return out
}

func main() {
	c := gen(2, 3)
	out := sq(c)

	fmt.Println(<-out)
	fmt.Println(<-out)

	// 因为 sq 的输入和输出是相同的，所以我们可以把 sq 嵌套使用
	for n := range sq(sq(gen(2, 3))) {
		fmt.Println(n)
	}
}
