package forthchapter

import "testing"

func TestArraySum(t *testing.T) {
	array := []int{1, 2, 7, 4, 5, 10, 6, 9, 8, 7, 3}
	t.Log(ArraySum(array))
}
