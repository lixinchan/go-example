package chapterone_test

import (
	"testing"
	. "chapterone"
)

func TestArrayStack_Push(t *testing.T) {
	var stack ArrayStack
	t.Log(stack)
	t.Log(stack.Size())
	for idx := 0; idx < 10; idx++ {
		stack.Push(idx)
	}
	t.Log(stack)
	for idx := 10; idx > 0; idx-- {
		if !stack.IsEmpty() {
			t.Log(stack.Pop())
		}
	}
}
