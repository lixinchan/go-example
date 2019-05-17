package main

import "fmt"

func main() {
	i, j := 12, 24

	p := &i
	fmt.Println(p)
	fmt.Println(*p)
	*p = 15
	fmt.Println(i)

	p = &j
	fmt.Println(p)
	fmt.Println(*p)
	*p = *p / 3
	fmt.Println(*p)
	fmt.Println(j)
}
