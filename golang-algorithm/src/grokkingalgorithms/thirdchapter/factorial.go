package thirdchapter

func Factorial(n int) int {
	if n <= 0 {
		return -1
	}
	if n == 1 {
		return 1
	}
	return n * Factorial(n-1)
}
