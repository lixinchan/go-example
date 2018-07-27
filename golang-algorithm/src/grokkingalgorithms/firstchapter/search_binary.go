package firstchapter

// binary search impl by cycle
func BinarySearchCycle(array []int, key int) int {
	if array == nil || len(array) == 0 {
		return -1
	}

	lo := 0
	hi := len(array) - 1
	for lo <= hi {
		mid := (lo + hi) >> 1
		if key > array[mid] {
			lo = mid + 1
		} else if key < array[mid] {
			hi = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

// binary search impl by recursive
func BinarySearchRecursive(array []int, lo int, hi int, key int) int {
	if array == nil || len(array) == 0 {
		return -1
	}
	if lo > hi {
		return -1
	}

	mid := (lo + hi) >> 1
	if key < array[mid] {
		return BinarySearchRecursive(array, lo, mid-1, key)
	} else if key > array[mid] {
		return BinarySearchRecursive(array, mid+1, hi, key)
	} else {
		return mid
	}
}
