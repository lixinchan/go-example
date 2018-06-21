package main

import "fmt"

func main() {
	fmt.Println("counting...")
	for idx := 0; idx < 10; idx++ {
		defer fmt.Println(idx)
	}
	defer fmt.Println("done.")
}
