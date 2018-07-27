package chapterone_test

import (
	. "chapterone"
	"testing"
	"fmt"
)

func TestBinarySearch(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 5, 6, 7, 8}
	retVal := SearchCycle(array, 3)
	fmt.Println(retVal)
	fmt.Println(SearchRecursion(array, 4, 1, 7))
	retVal = SearchLastIndex(array, 5)
	fmt.Println("ret:", retVal)
}
