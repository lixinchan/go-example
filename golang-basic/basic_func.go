// basic_func
package main

import (
	"fmt"
)

func main() {

	// 匿名函数&闭包
	var j int = 5
	a := func() func() {
		var i int = 10
		return func() {
			fmt.Println("i, j: %d, %d\n", i, j)
		}
	}()
	a()
	j *= 2
	a()
}
