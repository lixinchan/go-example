// golang_closure
package basic

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

func Closure() {
	ret := counterFactory(0)
	fmt.Println("%dn", ret())
	fmt.Println("%dn", ret())
	fmt.Println("%dn", ret())
}
