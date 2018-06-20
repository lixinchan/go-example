package main

import (
	"math"
	"fmt"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g => %g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Println(pow(2, 3, 10), pow(2, 5, 20))
}
