// basic_variables
package main

import (
	"fmt"
	//	"os"
)

func main() {
	// variables declare
	// var v1 int
	//	var v2 string
	//	var v3 [10]int // array
	//	var v4 []int   // slice
	//	var v5 struct {
	//		f int
	//	}
	//	var v6 *int           // pointer
	//	var v7 map[string]int // map key type:string, value type:int
	//	var v8 func(a int) int

	//	var (
	//		v9  int
	//		v10 string
	//	)

	//	// variables init
	//	var v11 int = 10
	//	var v12 int
	//	v12 = 10
	//	var v13 = 12
	//	v14 := 14

	// variables assignment
	//	var v15 int
	//	v15 = 15
	//	v16, v17 := 16, 17
	//	fmt.Println(v16)
	//	fmt.Println(v17)

	// give variable value
	//	var v18 int
	//	v18 = 18
	//	fmt.Println(v18)

	// ananymous variables
	//	func getName() (firstName, lastName string)  {
	//		return "brian", "chen"
	//	}

	//	_,lastName :=getName()

	// const

	// literal
	//	fmt.Println(-12)

	//	const PI = 3.14
	//	fmt.Println(PI)

	//	const a, b, c = 3, 4, "abc"
	//	fmt.Println(a, b, c)

	// 常量定义的右值可以是编译期运行的常量表达式
	//	const mask = 1 << 3
	//	fmt.Println(mask)

	// 右值不能是运行期才能知道的表达式
	//	const env = os.Getenv("HOME")
	//	env := os.Getenv("HOME")
	//	fmt.Println(env)

	// 预定义常量 true/false/iota
	//	const (
	//		a = iota
	//		b = iota
	//		c
	//	)
	//	fmt.Println(a, b, c)

	// 枚举
	//	const (
	//		SUN = iota
	//		MON
	//		TUE
	//		WEND
	//		THU
	//		FRI
	//		SAT
	//		numberOfDay
	//	)
	//	fmt.Println(MON, TUE, FRI, numberOfDay)

	//类型
	//bool:true/false
	var flag bool = true
	fmt.Println(flag)

}
