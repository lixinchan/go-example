// sort_selection.go
package chaptertwo

import "fmt"

func SelectionSort(array []int) {
	if array == nil {
		fmt.Println("nil")
		return
	}
	length := len(array)
	if length == 0 {
		fmt.Println("[]")
		return
	}
	for idx := 0; idx < length; idx++ {
		min := idx
		for jdx := idx + 1; jdx < length; jdx++ {
			if Less(array[jdx], array[min]) {
				min = jdx
			}
		}
		Exchange(array, idx, min)
	}
}

// selection sort enhance
func SelectionSortEnhance(array []int) {
	if array == nil {
		fmt.Println("nil")
		return
	}
	length := len(array)
	if length == 0 {
		fmt.Println("[]")
		return
	}
	for idx := 0; idx < length; idx++ {
		key := array[idx]
		jdx := idx - 1
		for jdx >= 0 && Less(key, array[jdx]) {
			array[jdx+1] = array[jdx]
			jdx--
		}
		array[jdx+1] = key
	}
}
