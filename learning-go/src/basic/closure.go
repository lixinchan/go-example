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
	fmt.Printf("%dn", ret())
	fmt.Printf("%dn", ret())
	fmt.Printf("%dn", ret())

}
