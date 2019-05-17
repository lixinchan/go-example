package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for idx := range pow {
		pow[idx] = 1 << uint(idx)
	}

	for _, ele := range pow {
		fmt.Printf("%d\n", ele)
	}
}
