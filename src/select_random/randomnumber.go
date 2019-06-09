package select_random

func RandomNumberGen(count int) <-chan int {
	ch := make(chan int)

	go func() {
		for i := 0; i < count; i++ {
			select {
			case ch <- 0:
			case ch <- 1:
			}
		}
	}()

	return ch
}
