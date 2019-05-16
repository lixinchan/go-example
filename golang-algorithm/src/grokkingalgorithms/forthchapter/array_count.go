package forthchapter

// ArrayCount count array ele
func ArrayCount(array []int) int {
	if array == nil || len(array) == 0 {
		return 0
	}
	return 1 + ArrayCount(array[1:])
}
