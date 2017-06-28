// search_binary
package main

import (
	"fmt"
)

// binary search cycle impl
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

// binary search recursion impl
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

// search the key last present index
func search_last_index(array []int, key int) int {
	if array == nil {
		return -1
	}
	low := 0
	high := len(array) - 1
	for low < high {
		mid := low + ((high - low) >> 1) + ((high - low) & 1)
		fmt.Println(mid)
		if array[mid] > key {
			high = mid - 1
		} else {
			low = mid
		}
	}

	if array[low] == key {
		return low
	} else {
		return -1
	}
	return -1
}

func main() {
	var array []int = []int{1, 2, 3, 4, 5, 5, 6, 7, 8}
	//	retVal := search_cycle(array, 3)
	//	fmt.Println(retVal)
	//	fmt.Println(search_recursion(array, 4, 1, 7))
	retVal := search_last_index(array, 5)
	fmt.Println("ret:", retVal)
}
