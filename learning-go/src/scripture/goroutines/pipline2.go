package goroutines

import "fmt"

// Cal cal
func Cal() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for idx := 1; idx < 10; idx++ {
			naturals <- idx
		}
		close(naturals)
	}()

	go func() {
		for val := range naturals {
			squares <- val * val
		}
		close(squares)
	}()

	for ret := range squares {
		fmt.Println(ret)
	}
}
