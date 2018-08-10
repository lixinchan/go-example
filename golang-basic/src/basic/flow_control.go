// FlowControl
package basic

import (
	"fmt"
)

func exam(x int) bool {
	if x >= 0 {
		return true
	} else {
		return false
	}
}

func example(x int) int {
	if x == 0 {
		return 5
	} else {
		return x
	}
}

func myGoto() {
	i := 0
HERE:
	fmt.Println(i)
	i++
	if i < 10 {
		goto HERE
	}
}

func FlowControl() {
	//	var x int = -5

	// if test
	//	var ret bool = exam(x)
	//	fmt.Println("retVal:", ret)

	//	retVal := example(x)
	//	fmt.Println("retVal:", retVal)

	// switch case test
	//	switch {
	//	case x == 0 || x > 0:
	//		fmt.Println("Great than zero")
	//	case x < 0:
	//		fallthrough
	//	case x < 0 && x > -3:
	//		fmt.Println("Less than zero")
	//	}

	// for loop
	//	sum := 0
	//	for i := 0; i <= 100; i++ {
	//		sum += i
	//	}

	//	for {
	//		sum++
	//		if sum > 5050 {
	//			break
	//		}
	//	}

	//	fmt.Println("sum:", sum)

	//	a := []int{1, 2, 3, 4, 5, 6, 7}
	//	fmt.Println("Before exchange:", a)
	//	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
	//		a[i], a[j] = a[j], a[i]
	//	}
	//	fmt.Println("After exchange:", a)
	myGoto()
}
