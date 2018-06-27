package main

import "fmt"

func main() {
	src := []int{1, 2, 3}
	fmt.Println("src:", src)
	dst := [] int{4, 5}
	fmt.Println("dst:", dst)
	copy(dst, src)
	fmt.Println("src:", src)
	fmt.Println("dst:", dst)
}
