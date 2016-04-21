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
}
