package thirdchapter

func FactorialTail(n, val int) int {
	if n == 1 {
		return val
	}
	return FactorialTail(n-1, val*n)
}
