package common

import "testing"

func TestLess(t *testing.T) {
	t.Log(Less(2, 3))
}

func TestExchange(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 5, 6, 7, 8}
	PrintArray(array)
	Exchange(array, 2, 3)
	PrintArray(array)
}

func TestPrintArray(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 5, 6, 7, 8}
	PrintArray(array)
}
