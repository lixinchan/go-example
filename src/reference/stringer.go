package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) String() string {
	return fmt.Sprintf("(%v's age is %d)", p.name, p.age)
}

func main() {
	p := person{"clx", 27}
	fmt.Println(p)
}
