package main

import "fmt"

func main() {
	sum := 1
	for ; sum < 5000; {
		sum += sum
	}
	fmt.Println("sum:", sum)
}
