package chapterone_test

import (
	. "chapterone"
	"testing"
)

// test fibonacci impl by tail recursive
func TestFibonacciTail(t *testing.T) {
	t.Log(FibonacciTail(150, 0, 1))
}

func BenchmarkFibonacciTail(b *testing.B) {
	b.Log(FibonacciTail(150, 0, 1))
}
