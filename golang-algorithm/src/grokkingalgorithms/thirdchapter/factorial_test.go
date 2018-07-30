package thirdchapter

import "testing"

func TestFactorial(t *testing.T) {
	t.Log(Factorial(20))
}

func BenchmarkFactorial(b *testing.B) {
	b.Log(Factorial(25))
}
