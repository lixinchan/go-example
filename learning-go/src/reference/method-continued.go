package main

import "fmt"

type selfFloat float64
type selfInt int

func (f selfFloat) abs() float64 {
	if f < 0 {
		return float64(-f)
	} else {
		return float64(f)
	}
}

func (i selfInt) abs() int {
	if i < 0 {
		return int(-i)
	} else {
		return int(i)
	}
}

func main() {
	f := selfFloat(-5)
	fmt.Println(f.abs())

	i := selfInt(-0)
	fmt.Println(i.abs())
}
