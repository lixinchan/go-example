package main

import "fmt"

func main() {
	sum := 0
	for idx := 0; idx < 10; idx++ {
		sum += idx
	}
	fmt.Println("sum:", sum)
}
