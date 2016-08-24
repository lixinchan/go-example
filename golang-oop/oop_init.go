// oop_init.go init
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

// return new rect object
func NewRect(x, y, weight, height float64) *Rect {
	return &Rect{x, y, weight, height}
}

func main() {
	// init Rect
	rect1 := new(Rect)
	rect2 := &Rect{}
	rect3 := &Rect{0, 0, 100, 200}
	rect4 := &Rect{weight: 100, height: 200}

	fmt.Println("Hello World!")
}
