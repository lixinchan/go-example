package main

import "fmt"

func fibonacciForChan(n int, ch chan int) {
	x, y := 0, 1
	for idx := 0; idx < n; idx++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func main() {
	ch := make(chan int, 10)
	go fibonacciForChan(cap(ch), ch)
	for val := range ch {
		fmt.Println(val)
	}
}
