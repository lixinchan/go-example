// search_binary
package main

import (
	"fmt"
)

func search(array []int, key int) int {
	low := 0
	high := len(array) - 1
	for low <= high {
		mid := low + (high-low)/2
		if key > array[mid] {
			low = mid + 1
		} else if key < array[mid] {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	var array []int = []int{1, 2, 3, 4, 5, 6, 7}
	retVal := search(array, 3)
	fmt.Println(retVal)
}
