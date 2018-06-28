package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for idx := 0; idx < 10; idx++ {
		fmt.Println(pos(idx), neg(-2*idx))
	}
}
