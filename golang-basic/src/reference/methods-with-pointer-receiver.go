package main

import (
	"math"
	"fmt"
)

type matrixForMethodPointer struct {
	x, y float64
}

func (m *matrixForMethodPointer) scale(f float64) {
	m.x = m.x * f
	m.y = m.y * f
}

func (m *matrixForMethodPointer) abs() float64 {
	return math.Sqrt(m.x*m.x + m.y*m.y)
}

func main() {
	v := &matrixForMethodPointer{3, 4}
	fmt.Printf("before scaling: %+v, abs: %v\n", v, v.abs())
	v.scale(5)
	fmt.Printf("after scaling: %+v, abs: %v\n", v, v.abs())

}
