package main

import "fmt"

func descInter(inter interface{}) {
	fmt.Printf("(%v, %T)\n", inter, inter)
}

func main() {
	var inter interface{}
	descInter(inter)

	inter = 21
	descInter(inter)

	inter = "hello"
	descInter(inter)
}
