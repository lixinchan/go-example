package main

import "fmt"

// return sum result
func sum(array [] int, ch chan int) {
	sum := 0
	for _, val := range array {
		sum += val
	}
	ch <- sum
}

func main() {
	array := []int{1, 3, 5, 7, 9, 11}

	ch := make(chan int)

	go sum(array[:len(array)/2], ch)
	go sum(array[len(array)/2:], ch)

	x, y := <-ch, <-ch
	fmt.Println(x, y, x+y)
}
