// Integer.go 为类型添加方法
package main

import (
	"fmt"
)

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

func main() {
	var a Integer = 3
	val := a.Less(5)
	fmt.Println(val)
}
