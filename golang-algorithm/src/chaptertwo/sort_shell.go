// sort_shell
package chaptertwo

// shell sort
func ShellSort(array []int) {
	if array == nil {
		return
	}
	if len(array) == 0 {
		return
	}
	length := len(array)
	h := 1
	for ; h < length/3; {
		h = 3*h + 1
	}
	for ; h >= 1; {
		for idx := h; idx < length; idx++ {
			for jdx := idx; jdx > h && Less(array[jdx], array[jdx-h]); jdx -= h {
				Exchange(array, jdx, jdx-h)
			}
		}
		h = h / 3
	}
}
