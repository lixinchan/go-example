package basic

import (
	"fmt"
)

func multiRetVal(n int) (ret int, err error) {
	return
}

func namedRetVal(n int) (ret int) {
	fmt.Println(ret)
	ret = n * n
	fmt.Println(ret)
	return
}

func deferFunc() {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d", i)
	}
}

func trace(str string) string {
	fmt.Println("entering:", str)
	return str
}

func untrace(str string) {
	fmt.Println("leaving:", str)
}

func deferFirst() {
	defer untrace(trace("first"))
	fmt.Println("in first")
}

func deferSecond() {
	defer untrace(trace("second"))
	fmt.Println("in second")
	deferFirst()
}

func Function() {
	//	fmt.Println(namedRetVal(2))
	//	for i := 0; i < 5; i++ {
	//		fmt.Print(i)
	//		defer fmt.Printf("%d", i)
	//	}
	//	defer deferFunc()
	deferSecond()
}
