// sort_selection.go
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

func selectionSort(array []int) {
	var N int = len(array)
	for i := 0; i < N; i++ {
		var min int = i
		for j := i + 1; j < N; j++ {
			if less(array[j], array[min]) {
				min = j
			}
		}
		exch(array, i, min)
	}
}

func main() {
	var array []int = []int{5, 7, 6, 1, 4, 3, 2, 0}
	printArray(array)
	selectionSort(array)
	printArray(array)
}
