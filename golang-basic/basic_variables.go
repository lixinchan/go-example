// basic_variables
package main

import (
	"fmt"
)

func main() {
	// variables declare
	var v1 int
	var v2 string
	var v3 [10]int // array
	var v4 []int   // slice
	var v5 struct {
		f int
	}
	var v6 *int           // pointer
	var v7 map[string]int // map key type:string, value type:int
	var v8 func(a int) int

	var (
		v9  int
		v10 string
	)

	// variables init
	var v11 int = 10
	var v12 int
	v12 = 10
	var v13 = 12
	v14 := 14

	// variables assignment
	var v15 int
	v15 = 15
	v16, v17 := 16, 17
	fmt.Println(v16)
	fmt.Println(v17)
}
