package main

import "fmt"

type nilInterface interface {
	testNil()
}

type testStrNil struct {
	s string
}

func (t *testStrNil) testNil() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.s)
}

func descNil(inter nilInterface) {
	fmt.Printf("(%v, %T)\n", inter, inter)
}

func main() {
	var inter nilInterface

	var t *testStrNil
	inter = t
	descNil(inter)
	inter.testNil()

	inter = &testStrNil{"hello"}
	descNil(inter)
	inter.testNil()
}
