package main

import (
	"math"
	"fmt"
)

type matrixs struct {
	x, y float64
}

func (m matrixs) abs() float64 {
	return math.Sqrt(m.x*m.x + m.y*m.y)
}

func (m *matrixs) scale(f float64) {
	m.x = m.x * f
	m.y = m.y * f
}

func main() {
	m := matrixs{3, 4}
	m.scale(10)
	fmt.Println(m.abs())
}
