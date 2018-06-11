package main

import (
	"fmt"
	"math"
)

func calculate() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64((x*x + y*y)))
	var u uint = uint(f)
	fmt.Println(x, y, f, u)
}

func main() {
	var i int = 32
	var f float64 = float64(i)
	var u uint = uint(f)

	fmt.Printf("Type: %T value:%v\n", i, i)
	fmt.Printf("Type: %T value:%v\n", f, f)
	fmt.Printf("Type: %T value:%v\n", u, u)

	calculate()
}
