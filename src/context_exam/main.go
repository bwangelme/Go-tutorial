package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(num int, wg *sync.WaitGroup, ctx context.Context) {
	defer func() {
		fmt.Printf("End %d\n", num)
		wg.Done()
	}()

	fmt.Printf("Start %d\n", num)
	for {
		select {
		default:
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(i, &wg, ctx)
	}

	time.Sleep(time.Second * 1) // 等待所有的 Goroutine 都执行过一遍 select 的 default
	cancel()
	wg.Wait()
	fmt.Println("All exit")
}
