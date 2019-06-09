package main

import (
	"context"
	"fmt"
)

func GenNumber(ctx context.Context) <-chan int {
	ch := make(chan int)

	go func() {
		for i := 2; ; i++ {
			select {
			case <-ctx.Done():
				return
			case ch <- i:
			}
		}
	}()

	return ch
}

func PrimeFilter(ctx context.Context, in <-chan int, prime int) <-chan int {
	ch := make(chan int)

	go func() {
		for {
			if i := <-in; i%prime != 0 {
				select {
				case <-ctx.Done():
					return
				case ch <- i:
				}
			}
		}
	}()

	return ch
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	ch := GenNumber(ctx)
	for i := 0; i < 1000; i++ {
		prime := <-ch
		fmt.Printf("%d: %d\n", i+1, prime)
		ch = PrimeFilter(ctx, ch, prime) // 基于出现的素数构造新的素数筛
	}

	cancel()
}
