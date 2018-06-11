package main

import "fmt"

func main() {
	var r rune
	var b bool
	var s string
	var c complex128
	fmt.Printf("Type: %T default value:%v\n", r, r)
	fmt.Printf("Type: %T default value:%v\n", b, b)
	fmt.Printf("Type: %T default value:%v\n", s, s)
	fmt.Printf("Type: %T default value:%v\n", c, c)
}
