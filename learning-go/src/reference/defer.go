package main

import "fmt"

func main() {
	// defer 会推迟执行函数直至外层函数执行完，但被推迟执行的函数的参数会先行执行
	defer fmt.Println(testDefer(test()))
	fmt.Println("hello,world.")
	fmt.Println("hahahah.")
}

func testDefer(idx int) int {
	return idx
}

func test() int {
	fmt.Println("calculating...")
	return 1 + 5
}
