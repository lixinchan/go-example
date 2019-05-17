package main

import "fmt"

type nilInter interface {
	nilValue()
}

func descNilVal(n nilInter) {
	fmt.Printf("(%v, %T\n)", n, n)
}

func main() {
	var inter nilInter
	descNilVal(inter)
	inter.nilValue()
}
