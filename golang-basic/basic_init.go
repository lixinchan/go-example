// basic_init.go

package main

import (
	"fmt"
)

func main() {
	// enum
	type ByteSize float64
	const (
		KB ByteSize = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
	)
	fmt.Println(KB)
	fmt.Println(GB)
}
