package main

import (
	"math"
	"fmt"
)

type matrix struct {
	x, y float64
}

func (m matrix) sqrt() float64 {
	return math.Sqrt(m.x*m.x + m.y*m.y)
}

func (m matrix) abs() (x, y float64) {
	return math.Abs(m.x), math.Abs(m.y)
}

func main() {
	m := matrix{3, 4}
	fmt.Println(m.sqrt())
	fmt.Println(m.abs())
}
