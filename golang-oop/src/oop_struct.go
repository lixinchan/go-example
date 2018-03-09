// struct.go 结构体
package main

import (
	"fmt"
)

type Rect struct {
	x, y           float64
	weight, height float64
}

func (r *Rect) Area() float64 {
	return r.weight * r.height
}

func main() {
	fmt.Println("Hello World!")
}
