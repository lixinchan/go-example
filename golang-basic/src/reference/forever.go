package main

import "fmt"

func main() {
	sum := 1
	for {
		sum += sum
		if sum > 1024 {
			break
		}
	}
	fmt.Println("sum:", sum)
}
