package main

import (
	"math"
	"fmt"
)

type abser interface {
	abs() float64
}

type iFloat float64

type matrixForInterface struct {
	x, y float64
}

func (f iFloat) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (m *matrixForInterface) abs() float64 {
	return math.Sqrt(m.x*m.x + m.y*m.y)
}

func main() {
	var ab abser
	f := iFloat(-math.Sqrt2)
	m := matrixForInterface{3, 4}

	ab = f
	fmt.Println(ab.abs())
	ab = &m
	fmt.Println(ab.abs())

}
