package forthchapter

// sum of array
func ArraySum(array []int) int {
	if array == nil {
		return 0
	}
	sum := 0
	for _, val := range array {
		sum += val
	}
	return sum
}
