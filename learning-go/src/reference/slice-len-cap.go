package main

import "fmt"

func main() {
	s := []int{1, 3, 5, 7, 9, 11, 13, 15}
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[2:]
	printSlice(s)

	s = s[0:5]
	printSlice(s)

	o := make([]int, 1, 10)
	printSlice(o)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
