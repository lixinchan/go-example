package main

import "fmt"

type vex struct {
	x int
	y int
}

func main() {
	v := vex{2, 3}
	fmt.Println(v.x)
	fmt.Println(v.y)
}
