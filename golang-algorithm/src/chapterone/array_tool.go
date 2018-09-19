package chapterone

// inplace swap
func inplaceSwap(x, y *int) {
	*y = *x ^ *y
	*x = *x ^ *y
	*y = *x ^ *y
}

// reverse array
func ReverseArray(array []int, cnt int) {
	for first, last := 0, cnt-1; first < last; first, last = first+1, last-1 {
		inplaceSwap(&array[first], &array[last])
	}
}
