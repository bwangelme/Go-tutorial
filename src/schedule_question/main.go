package main

import "fmt"
import "time"
import "runtime"

func main() {
	var x int
	threads := runtime.GOMAXPROCS(0)
	for i := 0; i < threads; i++ {
		go func() {
			for {
				time.Sleep(0)
				x++
				runtime.Gosched()
			}
		}()
	}
	time.Sleep(time.Second * 1)
	fmt.Println("x =", x)
}
