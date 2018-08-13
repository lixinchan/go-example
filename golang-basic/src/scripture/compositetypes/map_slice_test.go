package compositetypes

import "testing"

func TestCount(t *testing.T) {
	array := []string{"c", "l", "x"}
	Add(array)
	t.Log(Count(array))
}
