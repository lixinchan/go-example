package main

import "fmt"

func initChan(ch chan int) {
	for idx := 0; idx < 3; idx++ {
		ch <- idx
	}
	close(ch)
}
func main() {
	ch := make(chan int, 3)
	initChan(ch)
	for val := range ch {
		fmt.Println(val)
	}
}
