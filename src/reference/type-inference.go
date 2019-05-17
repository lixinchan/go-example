package main

import "fmt"

func main() {
	var i = 3.1
	// := 只能用于函数中
	j := 4 + 2i
	fmt.Printf("v is type of:%T\n", i)
	fmt.Printf("v is type of:%T\n", j)
}
