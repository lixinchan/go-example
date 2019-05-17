package main

import (
	"fmt"
)

type negativeErr float64

func (err negativeErr) Error() string {
	return fmt.Sprintf("Error number to sqrt:%v\n", err)
}

func sqrtErr(x float64) (float64, error) {
	return 0, nil
}

func main() {
	//fmt.Println(sqrtErr(2))
	//fmt.Println(sqrtErr(-2))
}
