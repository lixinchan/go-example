package main

import (
	"math"
	"fmt"
)

type matrixForRedirectContinued struct {
	x, y float64
}

func (m matrixForRedirectContinued) abs() float64 {
	return math.Sqrt(m.x*m.x + m.y*m.y)
}

func absFunc(m matrixForRedirectContinued) float64 {
	return math.Sqrt(m.x*m.x + m.y*m.y)
}

func main() {
	v := matrixForRedirectContinued{3, 4}
	fmt.Println(v.abs())
	fmt.Println(absFunc(v))

	p := &matrixForRedirectContinued{5, 12}
	fmt.Println(p.abs())
	fmt.Println((*p).abs())
	fmt.Println(absFunc(*p))
}
