package basic

import (
	"fmt"
)

func Question() {
	fizzBuzz(100)
}

func fizzBuzz(n int) {
	for idx := 1; idx <= n; idx++ {
		if idx%3 == 0 {
			if idx%5 == 0 {
				fmt.Println("FizzBuzz")
			} else {
				fmt.Println("Fizz")
			}
		} else if idx%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(idx)
		}
	}
}
