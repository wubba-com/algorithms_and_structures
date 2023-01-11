package utils

import "sync"

func Download(sizes ...int) <-chan int {
	var c = make(chan int, 2)
	go func() {
		wg := sync.WaitGroup{}
		for _, size := range sizes {
			wg.Add(1)
			go func(size int) {
				defer wg.Done()

				var sum int
				for i := 0; i <= size; i++ {
					sum += i
				}
				c <- sum

			}(size)
		}
		wg.Wait()
		close(c)
	}()
	return c
}
