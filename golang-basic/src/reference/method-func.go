package main

import (
	"math"
	"fmt"
)

type vetrix struct {
	x, y float64
}

func sqrts(v vetrix) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main() {
	v := vetrix{5, 12}
	fmt.Println(sqrts(v))
}
