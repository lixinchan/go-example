package main

import "fmt"

func main() {
	var inter interface{}
	inter = "hello"

	s := inter.(string)
	fmt.Println(s)

	s, ok := inter.(string)
	if ok {
		fmt.Println(s)
	}

	v, ok := inter.(int)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("panic...")
	}

	v = inter.(int)
	fmt.Println(v)
}
