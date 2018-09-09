package chapterone_test

import (
	. "chapterone"
	"testing"
)

func TestLinkedStack_Push(t *testing.T) {
	var stack LinkedStack
	node := InitLinkedStack()
	stack = node
	t.Log(stack.Size())
	for idx := 0; idx < 10; idx++ {
		stack.Push(idx)
	}
	t.Log(stack.Size())
	size := stack.Size()
	for idx := 0; idx < size; idx++ {
		t.Log(stack.Top())
		t.Log(stack.Pop())
	}
	t.Log(stack.Clear())
	t.Log(stack.Size())
}
