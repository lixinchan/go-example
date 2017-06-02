// shellsort.go
package shellsort

func ShellSort(array []int) {
	var length int = len(array)
	var h int = 1
	if h < length/3 {
		h = 3*h + 1
	}
	if h >= 1 {
		for idx := h; idx < length; idx++ {
			for j := idx; j >= h && less(array[j], array[j-h]); j -= h {
				exch(array, j, j-h)
			}
		}
		h = h / 3
	}
}

func less(v int, w int) bool {
	return v-w < 0
}

func exch(array []int, i int, j int) {
	array[i], array[j] = array[j], array[i]
}

func printArray(array []int) {
	for _, val := range array {
		fmt.Print(val, " ")
	}
	fmt.Println()
}
