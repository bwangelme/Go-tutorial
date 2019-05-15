package main

import (
	"fmt"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

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

func merge(cs ...<-chan int) <-chan int {
	// 为 channel 加上 Buffer，不再需要创建额外的 Goroutine，直接使用当前函数将输入写入到一个会关闭的 Buffer 中即可。
	// 但这样的代码存在的问题是 channel 的 Buffer 定位7是因为我们知道数据的个数，如果不知道数据的个数，channel 的 Buffer 就无法确定
	out := make(chan int, 7)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
	}

	for _, c := range cs {
		output(c)
	}
	close(out)

	return out
}

func main() {
	in := gen(2, 3, 33, 1, 32, 17, 55)

	// 将 in 中的数据分别分发到两个 Goroutine 中
	c1 := sq(in)
	c2 := sq(in)

	// 将 c1 和 c2 中的数据消费并合并到一起
	for n := range merge(c1, c2) {
		fmt.Println(n) // 4 then 9, or 9 then 4
	}
}
