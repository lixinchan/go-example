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
	//	var array [5]int = [5]int{0, 1, 2, 3, 4}
	//	var mySlice []int = array[1:5]
	//	for _, v := range array {
	//		fmt.Println("ori array:", v)
	//	}
	//	for _, v := range mySlice {
	//		fmt.Println("slice:", v)
	//	}

	// declare array slice directly
	// 5 element
	//	mySlice1 := make([]int, 5)
	// 5 element & 10 cap
	//	mySlice2 := make([]int, 5, 10)
	// init directly
	//	mySlice3 := []int{1, 2, 3, 4, 5}

	// len() and cap() func
	//	fmt.Println("len of mySlice2:", len(mySlice2))
	//	fmt.Println("cap of mySlice2:", cap(mySlice2))

	// append() func
	//	mySlice3 = append(mySlice3, 6, 7, 8)
	//	mySlice4 := []int{9, 10}
	//	mySlice3 = append(mySlice3, mySlice4...)
	//	fmt.Println(mySlice3)
	// based on slice
	//	newSlice := mySlice3[3:8]
	//	fmt.Println(newSlice)

	// copy() func
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{6, 7, 8}
	// copy(slice1, slice2)
	copy(slice2, slice1)
	fmt.Println(slice1)
	fmt.Println(slice2)
}
