// search_binary
package main

import (
	"fmt"
)

// binary search cycle impl
func searchCycle(array []int, key int) int {
	if array == nil || len(array) == 0 {
		return -1
	}

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
func searchRecursion(array []int, key int, low int, high int) int {
	if array == nil || len(array) == 0 {
		return -1
	}
	if low > high {
		return -1
	}
	mid := low + (high-low)/2
	if key > array[mid] {
		return searchRecursion(array, key, mid+1, high)
	} else if key < array[mid] {
		return searchRecursion(array, key, low, mid-1)
	} else {
		return mid
	}
}

// search the key last present index
func searchLastIndex(array []int, key int) int {
	if array == nil || len(array) == 0 {
		return -1
	}
	low := 0
	high := len(array) - 1
	for low < high {
		mid := low + ((high - low) >> 1) + ((high - low) & 1)
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
	array := []int{1, 2, 3, 4, 5, 5, 6, 7, 8}
	retVal := searchCycle(array, 3)
	fmt.Println(retVal)
	fmt.Println(searchRecursion(array, 4, 1, 7))
	retVal = searchLastIndex(array, 5)
	fmt.Println("ret:", retVal)
}
