package secondchapter

import "grokkingalgorithms/common"

// selection sort
func SelectionSort(array []int) {
	if array == nil || len(array) == 0 {
		return
	}

	minIdx := 0
	for idx := range array {
		minIdx = idx
		for jdx := idx + 1; jdx < len(array); jdx++ {
			if common.Less(array[jdx], array[minIdx]) {
				minIdx = jdx
			}
		}
		common.Exchange(array, idx, minIdx)
	}
}
