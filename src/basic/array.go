package basic

import (
	"fmt"
)

func Modify(array [5]int) {
	array[0] = 12
	fmt.Println("In modify:", array)
}

func ReverseArray(array []int) []int {
	length := len(array)
	for i := 0; i < length/2; i++ {
		temp := array[i]
		array[i] = array[length-1-i]
		array[length-i-1] = temp
	}
	return array
}

func Array() {
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
	// array := [5]int{0, 1, 2, 3, 4}
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

	//	 declare array slice directly
	//	 5 element
	//	mySlice1 := make([]int, 5)
	//	 5 element & 10 cap
	//	mySlice2 := make([]int, 5, 10)
	//	 init directly
	//	mySlice3 := []int{1, 2, 3, 4, 5}

	//	 len() and cap() func
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
	slice1 := []int{1, 2, 3}
	slice2 := []int{6, 7, 8, 9}
	//	copy(slice1, slice2)
	copy(slice2, slice1)
	fmt.Println(slice1)
	fmt.Println(slice2)

	// reverse array
	//	var array []int = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	//	fmt.Println(array)
	//	retArray := reverseArray(array)
	//	fmt.Println(retArray)

	//arr := [...]int{1, 2, 3}
	//b := &arr

	//fmt.Println(arr[1], arr[2])
	//fmt.Println(b[1], b[2])
	//for idx, v := range b {
	//	fmt.Println(idx, v)
	//}

	var times [5][0]int
	for idx := range times {
		fmt.Println(idx, "hello")
	}

	for range times {
		fmt.Println("hello")
	}

}
