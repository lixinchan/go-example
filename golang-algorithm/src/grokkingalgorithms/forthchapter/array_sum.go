package forthchapter

// ArrayLoopSum sum of array loop
func ArrayLoopSum(array []int) int {
	if array == nil {
		return 0
	}
	sum := 0
	for _, val := range array {
		sum += val
	}
	return sum
}

// ArrayRecursiveSum sum of array recursive
func ArrayRecursiveSum(array []int) int {
	if len(array) == 0 {
		return 0
	}
	return array[0] + ArrayRecursiveSum(array[1:])
}
