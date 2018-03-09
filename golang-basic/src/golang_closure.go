// golang_closure
package main

import (
	"fmt"
)

func counterFactory(j int) func() int {
	i := j
	return func() int {
		i++
		return i
	}
}

func main() {
	ret := counterFactory(0)
	fmt.Println("%dn", ret())
	fmt.Println("%dn", ret())
	fmt.Println("%dn", ret())
}
