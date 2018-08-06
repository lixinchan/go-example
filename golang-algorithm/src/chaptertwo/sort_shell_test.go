package chaptertwo_test

import (
	. "chaptertwo"
	"testing"
)

func TestShellSort(t *testing.T) {
	array := []int{1, 2, 7, 4, 5, 10, 6, 9, 8, 7, 3}
	PrintArray(array)
	ShellSort(array)
	t.Log(array)
}
