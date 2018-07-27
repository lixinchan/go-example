// sort_selection.go
package chaptertwo

func SelectionSort(array []int) {
	length := len(array)
	for i := 0; i < length; i++ {
		var min int = i
		for j := i + 1; j < length; j++ {
			if Less(array[j], array[min]) {
				min = j
			}
		}
		Exchange(array, i, min)
	}
}
