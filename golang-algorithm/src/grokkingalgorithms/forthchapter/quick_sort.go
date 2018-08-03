package forthchapter

// quick sort
func QuickSort(array []int) []int {
	if array == nil {
		return nil
	}
	if len(array) < 2 {
		return array
	} else {
		pivot := array[0]
		var less = []int{}
		var grater = []int{}
		for _, val := range array[1:] {
			if val <= pivot {
				less = append(less, val)
			} else {
				grater = append(grater, val)
			}
		}
		less = append(QuickSort(less), pivot)
		grater = QuickSort(grater)
		return append(less, grater...)
	}
}
