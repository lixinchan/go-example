package main

import "fmt"

type interf interface {
	testInter()
}

type vFloat float64

type testStr struct {
	s string
}

func (t *testStr) testInter() {
	fmt.Println(t.s)
}

func (v vFloat) testInter() {
	fmt.Println(v)
}

func desc(inter interf) {
	fmt.Printf("%v %T\n", inter, inter)
}

func main() {
	var inter interf
	s := &testStr{"hello"}
	inter = s
	desc(inter)
	inter.testInter()

	v := vFloat(2)
	inter = v
	desc(inter)
	inter.testInter()
}
