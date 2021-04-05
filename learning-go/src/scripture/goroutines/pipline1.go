package goroutines

import "fmt"

// Calculator cal
func Calculator() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for idx := 1; idx < 10; idx++ {
			naturals <- idx
		}
	}()

	go func() {
		for {
			x := <-naturals
			squares <- (x * x)
		}
	}()

	for {
		fmt.Println(<-squares)
	}
}
