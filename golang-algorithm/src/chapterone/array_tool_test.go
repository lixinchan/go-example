package chapterone_test

import (
	. "chapterone"
	"testing"
)

func TestReverseArray(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	ReverseArray(array, len(array))
	t.Log(array)
}
