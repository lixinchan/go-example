package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlices("a", a)

	b := make([]int, 0, 5)
	printSlices("b", b)

	c := b[:2]
	printSlices("c", c)

	d := b[2:5]
	printSlices("d", d)
}

func printSlices(s string, slice []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(slice), cap(slice), slice)
}
