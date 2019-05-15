package main

import (
	"fmt"
	"sync"

	"github.com/temprory/util"
)

func main() {
	wg := sync.WaitGroup{}
	cl := util.NewCorsLink("test", 10, 3)
	for i := 0; i < 30; i++ {
		idx := i
		wg.Add(1)
		cl.Go(func(task *util.LinkTask) {
			defer wg.Done()

			task.WaitPre()
			if idx%3 == 0 {
				fmt.Println("A")
			} else if idx%3 == 1 {
				fmt.Println("B")
			} else {
				fmt.Println("C")
			}

			task.Done(nil)
		})
	}
	wg.Wait()
	cl.Stop()
}
