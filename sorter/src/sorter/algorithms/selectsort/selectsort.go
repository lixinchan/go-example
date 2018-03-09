// selectsort.go
package selectsort

import (
	"fmt"
)

func SelectSort(array []int) {
	var length int := len(array)
	for i := 0; i < length; i++ {
		var min int = i
		for j := i + 1; j < length; j++ {
			if less(array[j], array[min]) {
				min = j
			}
		}
		exch(array, i, min)
	}
}

func less(v int, w int) bool {
	return v-w < 0
}

func exch(array []int, i int, j int) {
	array[i], array[j] = array[j], array[i]
}

func printArray(array []int) {
	for _, val := range array {
		fmt.Print(val, " ")
	}
	fmt.Println()
}
