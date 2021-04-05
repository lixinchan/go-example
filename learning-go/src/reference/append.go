package main

import "fmt"

func main() {
	var slice []int
	printAllSlice(slice)

	slice = append(slice, 1)
	printAllSlice(slice)

	slice = append(slice, 2, 3, 4, 5)
	printAllSlice(slice)

	s := []int{6, 7, 8}
	slice = append(slice, s...)
	printAllSlice(slice)
}

func printAllSlice(slice []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(slice), cap(slice), slice)
}
