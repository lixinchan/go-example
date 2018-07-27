package chaptertwo_test

import (
	. "chaptertwo"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	array := []int{5, 7, 6, 1, 4, 3, 2}
	PrintArray(array)
	SortEnhance(array)
	PrintArray(array)
}
