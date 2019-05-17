package main

import (
	"math"
	"fmt"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	//fmt.Println(pow(2, 3, 10))
	//fmt.Println(pow(2, 5, 20))
	fmt.Println(pow(2, 3, 10), pow(2, 5, 20))
}
