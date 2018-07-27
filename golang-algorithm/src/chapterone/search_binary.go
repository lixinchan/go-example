// search_binary
package chapterone

// binary search cycle impl
func SearchCycle(array []int, key int) int {
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
func SearchRecursion(array []int, key int, low int, high int) int {
	if array == nil || len(array) == 0 {
		return -1
	}
	if low > high {
		return -1
	}
	mid := low + (high-low)/2
	if key > array[mid] {
		return SearchRecursion(array, key, mid+1, high)
	} else if key < array[mid] {
		return SearchRecursion(array, key, low, mid-1)
	} else {
		return mid
	}
}

// search the key last present index
func SearchLastIndex(array []int, key int) int {
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
