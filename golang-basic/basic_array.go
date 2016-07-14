// basic_array.go
package main

import (
	"fmt"
)

func modify(array [5]int) {
	array[0] = 12
	fmt.Println("In modify:", array)
}

func main() {
	// declare
	// byte type
	//	var byteArr []byte = [32]byte{}
	// complex type
	//	[2*N] struct {x, y int}
	// pointer type
	//	[100]*float64
	// 二维
	//	[2][3]int

	//len()
	//	var array [5]int = [5]int{0, 1, 2, 3, 4}
	//	fmt.Println("The array's length:", len(array))

	// iter
	//	for i := 0; i < len(array); i++ {
	//		fmt.Println("array[", i, "]=", array[i])
	//	}

	// iter range
	//	for i, v := range array {
	//		fmt.Println("array[", i, "]=", v)
	//	}
	//	for _, v := range array {
	//		fmt.Println(v)
	//	}

	// array is value type can not modify array's value
	//	array := [5]int{0, 1, 2, 3, 4}
	//	modify(array)
	//	fmt.Println("In main:", array)

	// array slice
	// base on array

	// declare array slice directly
}
