package main

import "fmt"

func fibonacci() func() int {
	a, b := 0, 1
	sum := 0
	return func() int {
		if a == 0 && b == 1 {
			fmt.Println(a)
			fmt.Println(b)
		}
		sum = a + b
		a = b
		b = sum
		return sum
	}
}

func main() {
	fn := fibonacci()
	for idx := 0; idx < 10; idx++ {
		fmt.Println(fn())
	}
}
