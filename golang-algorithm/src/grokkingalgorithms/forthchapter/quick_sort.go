package forthchapter

import "grokkingalgorithms/common"

// QuickSort sort array
func QuickSort(array []int) []int {
	if array == nil {
		return nil
	}
	if len(array) == 0 {
		return nil
	}
	sort(array, 0, len(array)-1)
	return array
}

func sort(array []int, lo, hi int) {
	if hi <= lo {
		return
	}
	// partition array
	idx := partition(array, lo, hi)
	// sort left
	sort(array, lo, idx-1)
	// sort right
	sort(array, idx+1, hi)
}

func partition(array []int, lo, hi int) int {
	idx := lo + 1
	jdx := hi
	pivot := array[lo]
	for true {
		for ; common.Less(array[idx], pivot); idx++ {
			if idx == hi {
				break
			}
		}
		for ; common.Less(pivot, array[jdx]); jdx-- {
			if jdx == lo {
				break
			}
		}
		if idx >= jdx {
			break
		}
		common.Exchange(array, idx, jdx)
	}
	common.Exchange(array, lo, jdx)
	return jdx
}
