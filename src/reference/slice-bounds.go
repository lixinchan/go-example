package main

import "fmt"

func main() {
	s := []int{1, 3, 5, 7, 9, 11}
	//fmt.Println(s)

	s = s[1:4]
	fmt.Println(s)

	s = s[:3]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

	s = s[:]
	fmt.Println(s)

	s = s[:4]
	fmt.Println(s)
}
