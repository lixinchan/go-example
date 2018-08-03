package forthchapter

import (
	"testing"
	"grokkingalgorithms/common"
)

func TestQuickSort(t *testing.T) {
	array := []int{1, 2, 7, 4, 5, 10, 6, 9, 8, 7, 3}
	common.PrintArray(array)
	t.Log(QuickSort(array))
}

func BenchmarkQuickSort(b *testing.B) {

}
