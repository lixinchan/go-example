package secondchapter

import (
	"testing"
	"grokkingalgorithms/common"
)

func TestSelectionSort(t *testing.T) {
	array := []int{1, 2, 7, 4, 5, 10, 6, 9, 8, 7, 3}
	common.PrintArray(array)
	SelectionSort(array)
	common.PrintArray(array)
}
