// basic_practice.go
package basic

import (
	"fmt"
)

func Practice() {

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
	//	const (
	//		FIZZ = 3
	//		BUZZ = 5
	//	)

	//	for index := 1; index <= 100; index++ {
	//		flag := false
	//		if index%FIZZ == 0 {
	//			fmt.Print("Fizz")
	//			flag = true
	//		}
	//		if index%BUZZ == 0 {
	//			fmt.Print("Buzz")
	//			flag = true
	//		}
	//		if !flag {
	//			fmt.Print(index)
	//		}
	//		fmt.Println()
	//	}

	// := 什么时候可用？本次声明的变量与index在同一作用域中，如果index在外层已经声明，
	//	那么新声明的这个变量将是一个新的变量，并且，类型需要一致，本次声明中至少有一个变量是新声明的
	//	var index int = 5
	//	if true {
	//		index := 3
	//		fmt.Println(index)
	//	}
	//	fmt.Println(index)
	//	idx, index := 1, 3
	//	fmt.Println(idx, index)

	// ++, -- 是语句而不是表达式，平行赋值时，不能使用++/--
	array := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
	fmt.Println(array)

}
