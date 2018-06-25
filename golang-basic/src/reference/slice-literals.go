package main

import "fmt"

func main() {
	o := [2]int{1, 2}
	fmt.Println(o)
	n := o[0:]
	fmt.Println(n)

	p := []bool{true, false, false, true}
	fmt.Println("p:", p)

	q := []int{1, 3, 5, 7, 9}
	fmt.Println("q:", q)

	r := []struct {
		i int
		b bool
	}{{1, false}, {2, true}, {3, false}}
	fmt.Println("r:", r)
}
