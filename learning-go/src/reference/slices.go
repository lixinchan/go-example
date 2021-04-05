package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11}

	var slice []int = primes[0:5]
	fmt.Println(slice)
	fmt.Println("len:", len(slice))
	fmt.Println("cap:", cap(slice))
}
