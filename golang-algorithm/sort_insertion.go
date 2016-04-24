// sort_insertion
package main

import (
	"fmt"
)

func less(v int, w int) bool {
	return v-w < 0
}

func exch(array []int, i int, j int) {
	array[i], array[j] = array[j], array[i]
}

func printArray(array []int) {
	for _, v := range array {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func sort(array []int) {
	var N int = len(array)
	for i := 0; i < N; i++ {
		for j := i; j > 0 && less(array[j], array[j-1]); j-- {
			exch(array, j, j-1)
		}
	}

}

func main() {
	var array []int = []int{5, 7, 6, 1, 4, 3, 2}
	printArray(array)
	sort(array)
	printArray(array)
}
