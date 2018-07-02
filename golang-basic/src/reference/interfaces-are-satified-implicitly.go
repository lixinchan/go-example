package main

import "fmt"

type inter interface {
	ab()
	bc()
}

type firstString struct {
	s string
}

type secondString struct {
	s string
}

func (str firstString) ab() {
	fmt.Println(str.s)
}

func (str firstString) bc() {
	fmt.Println(str.s)
}

func (str secondString) ab() {
	fmt.Println(str.s)
}

func (str secondString) bc() {
	fmt.Println(str.s)
}

func main() {
	var s inter
	a := firstString{"hello"}
	s = a
	s.ab()

	b := secondString{"world"}
	s = b
	s.bc()
}
