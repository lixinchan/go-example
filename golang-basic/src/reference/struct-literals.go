package main

import "fmt"

type vert struct {
	x, y int
}

var (
	v1 = vert{1, 2}
	v2 = vert{x: 5}
	v3 = vert{}
	p  = &vert{6, 7}
)

func main() {
	fmt.Println(v1, v2, v3, p)
}
