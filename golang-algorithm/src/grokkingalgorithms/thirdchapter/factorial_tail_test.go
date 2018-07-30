package thirdchapter

import "testing"

func TestFactorialTail(t *testing.T) {
	t.Log(FactorialTail(5, 1))
}

func BenchmarkFactorialTail(b *testing.B) {
	b.Log(FactorialTail(10, 1))
}
