package chaptertwo_test

import (
	"chaptertwo"
	"testing"
)

func TestSortByTree(t *testing.T) {
	array := []int{5, 7, 6, 1, 4, 3, 2}
	chaptertwo.SortByTree(array)
	chaptertwo.PrintArray(array)
}
