package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func addContinue(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(1, 3))
	fmt.Println(addContinue(1, 3))
}
