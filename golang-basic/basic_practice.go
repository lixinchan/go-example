// basic_practice.go
package main

import (
	"fmt"
)

func main() {

	// Q1
	//	for index := 0; index < 10; index++ {
	//		fmt.Println("index-", index, ":", index)
	//	}

	//	index := 0
	//flag:
	//	fmt.Println("index-", index, ":", index)
	//	if index < 9 {
	//		index++
	//		goto flag
	//	}

	//	array := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//	for k, v := range array {
	//		fmt.Println(k, ":", v)
	//	}

	// Q2 FizzBuzz
	for index := 1; index <= 100; index++ {
		switch index {
		case index%3 == 0:
			fmt.Println(index, "-Fizz")
		case index%5 == 0:
			fmt.Println(index, "-Buzz")
		case index%3 == 0 && index%5 == 0:
			fmt.Println(index, "-FizzBuzz")
		default:
			fmt.Println(index, "-", index)
		}
	}

}
