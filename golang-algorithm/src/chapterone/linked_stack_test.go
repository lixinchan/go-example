package chapterone_test

import (
	"testing"
	. "chapterone"
)

func TestLinkedStack_Push(t *testing.T) {
	var stack LinkedStack
	var node LinkedStackNode
	stack = &node

	t.Log(stack)
}
