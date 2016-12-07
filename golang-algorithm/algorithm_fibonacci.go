// algorithm_fibonacci
package main

import (
	"fmt"
)

func fibonacci(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 1
	}
	return fibonacci(N-1) + fibonacci(N-2)
}

func main() {
	for index := 0; index < 20; index++ {
		fmt.Println(index, ":", fibonacci(index))
	}
}
