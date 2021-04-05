package main

import "fmt"

func main() {
	m := make(map[string]int)
	fmt.Println("The value:", m)

	m["answer"] = 12
	fmt.Println("The value:", m["answer"])

	m["answer"] = 15
	fmt.Println("The value:", m["answer"])

	//delete(m, "answer")
	//fmt.Println("The value:", m["answer"])

	ele, ok := m["answer"]
	if ok {
		fmt.Println("The value:", ele)
	}
}
