// algorithm_fibonacci
package chapterone

func Fibonacci(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 1
	}
	return Fibonacci(N-1) + Fibonacci(N-2)
}
