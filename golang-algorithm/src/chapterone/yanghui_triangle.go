package chapterone

import (
	"fmt"
)

// print yanghui triangle
func PrintTriangle(line int) {
	if line <= 0 {
		return
	}
	array := make([]int, 10)
	for idx := 0; idx < line; idx++ {
		for jdx := 0; jdx < (line - idx); jdx++ {
			fmt.Print(" ")
		}
		for jdx := 0; jdx < (idx + 1); jdx++ {
			length := len(array)
			var value int
			if jdx == 0 || jdx == idx {
				value = 1
			} else {
				value = array[length-idx] + array[length-idx-1]
			}
			array = append(array, value)
			fmt.Print(value, " ")
		}
		fmt.Println()
	}
}
