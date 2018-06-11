package main

import (
	"math/cmplx"
	"fmt"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
	// uint8 alias
	Byte byte = 1
	// int32 alias
	Rune rune = 2
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Printf("Type: %T Value: %v\n", Byte, Byte)
	fmt.Printf("Type: %T Value: %v\n", Rune, Rune)
}
