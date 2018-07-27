package firstchapter

import (
	"testing"
)

// test binary search cycle
func TestBinarySearchCycle(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 5, 6, 7, 8}
	if ret := BinarySearchCycle(array, 5); ret != -1 {
		t.Log("Result:", ret)
	}
}

// test binary search recursive
func TestBinarySearchRecursive(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 5, 6, 7, 8}
	if ret := BinarySearchRecursive(array, 0, len(array)-1, 6); ret != -1 {
		t.Log("Result:", ret)
	}
}

// benchmark test binary search cycle
func BenchmarkBinarySearchCycle(b *testing.B) {

}

// benchmark test binary search recursive
func BenchmarkBinarySearchRecursive(b *testing.B) {

}
