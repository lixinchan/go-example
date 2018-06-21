package main

import "fmt"

type vetex struct {
	x int
	y int
}

func main() {
	v := vetex{2, 5}
	fmt.Println(v)
	p := &v
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println((*p).x)
	p.x = 1e9
	fmt.Println(v)

}
