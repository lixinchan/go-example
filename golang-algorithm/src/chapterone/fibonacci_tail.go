package chapterone

// fibonacci impl by tail recursive
func FibonacciTail(n, firstVal, secondVal int) int {
	if n == 0 {
		return firstVal
	}
	return FibonacciTail(n-1, secondVal, firstVal+secondVal)
}
