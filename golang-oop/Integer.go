// Integer.go 为类型添加方法
package main

import (
	"fmt"
)

type Integer int

// object oriented programing
func (a Integer) Less(b Integer) bool {
	return a < b
}

// processing oriented programing
func Integer_Less(a, b Integer) bool {
	return a < b
}

func (a Integer) Add(b Integer) {
	a += b
}

func (a *Integer) Add_Pointer(b Integer) {
	*a += b
}

func main() {
	var a Integer = 3
	val := a.Less(5)
	fmt.Println(val)
	fmt.Println("---")
	var ret bool = Integer_Less(5, 3)
	fmt.Println(ret)
	fmt.Println("---")
	a.Add(2)
	fmt.Println("add:", a)
	a.Add_Pointer(2)
	fmt.Println("add_pointer:", a)
}
