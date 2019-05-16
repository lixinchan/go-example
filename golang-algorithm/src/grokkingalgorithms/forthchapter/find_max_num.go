package forthchapter

import "grokkingalgorithms/common"

// FindMax find max array ele
func FindMax(array []int) int {
	if array == nil || len(array) == 0 {
		return -1
	}
	max := array[0]
	for idx := 0; idx < len(array); idx++ {
		if common.Less(max, array[idx]) {
			max = array[idx]
		}
	}
	return max
}
