package main

import (
	"math"
	"fmt"
)

type matrixForPointer struct {
	x, y float64
}

func abs(m matrixForPointer) float64 {
	return math.Sqrt(m.x*m.x + m.y*m.y)
}

func scale(m *matrixForPointer, f float64) {
	m.x = m.x * f
	m.y = m.y * f
}

func main() {
	m := matrixForPointer{3, 4}
	scale(&m, 10)
	fmt.Println(abs(m))
}
