// search_binary
package main

import (
	"fmt"
)

func search_cycle(array []int, key int) int {
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

func search_recursion(array []int, key int, low int, high int) int {
	if low > high {
		return -1
	}
	mid := low + (high-low)/2
	if key > array[mid] {
		return search_recursion(array, key, mid+1, high)
	} else if key < array[mid] {
		return search_recursion(array, key, low, mid-1)
	} else {
		return mid
	}
}

func main() {
	var array []int = []int{1, 2, 3, 4, 5, 6, 7}
	retVal := search_cycle(array, 3)
	fmt.Println(retVal)
	fmt.Println(search_recursion(array, 4, 1, 7))
}
