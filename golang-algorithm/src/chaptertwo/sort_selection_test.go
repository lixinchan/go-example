package chaptertwo_test

import (
	. "chaptertwo"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	array := []int{5, 7, 6, 1, 4, 3, 2, 0}
	PrintArray(array)
	SelectionSort(array)
	PrintArray(array)
}

func TestSelectionSortEnhance(t *testing.T) {
	array := []int{5, 7, 6, 1, 4, 3, 2, 0}
	PrintArray(array)
	SelectionSortEnhance(array)
	PrintArray(array)
}
