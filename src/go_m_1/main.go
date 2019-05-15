package main

// https://www.v2ex.com/t/556075#reply12
import (
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func(i int) {
			println(i)
			wg.Done()
		}(i)
	}

	for i := 10; i < 20; i++ {
		go func(i int) {
			println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
