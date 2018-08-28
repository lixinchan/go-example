package chapterone

import "errors"

// linkedStack's methods
type LinkedStack interface {
	Push(data interface{})     // push
	Pop() (interface{}, error) // pop
	Top() (interface{}, error) // top data, not remove
	Size() int                 // size
	IsEmpty() bool             // isEmpty
	Clear() bool               // clear stack
}

// linked stack
type LinkedStackNode struct {
	next *LinkedStackNode
	data interface{}
}

var (
	size   int              // size of linkedStack
	header *LinkedStackNode // header of linkedStack
)

// new entry
func createEntry(head *LinkedStackNode, data interface{}) *LinkedStackNode {
	return &LinkedStackNode{next: head, data: data}
}

// init linkedStackNode
func InitLinkedStack() *LinkedStackNode {
	header = createEntry(nil, nil)
	header.next = header
	return header
}

// push
func (stack *LinkedStackNode) Push(data interface{}) {
	oldHeader := header
	header = createEntry(oldHeader, data)
	size++
}

// pop
func (stack *LinkedStackNode) Pop() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("stack is nil")
	}
	popEntry := header
	header = popEntry.next
	popEntry.next = nil
	size--
	return popEntry.data, nil
}

// top
func (stack LinkedStackNode) Top() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("stack is nil")
	}
	return header.data, nil
}

// size
func (stack LinkedStackNode) Size() int {
	return size
}

// isEmpty
func (stack LinkedStackNode) IsEmpty() bool {
	return size == 0
}

func (stack *LinkedStackNode) Clear() bool {
	if stack.IsEmpty() {
		return true
	}
	for !stack.IsEmpty() {
		clearEntry := header
		header.next = clearEntry.next
		clearEntry.next = nil
		clearEntry.data = nil
		size--
	}
	header.next = header
	return true
}
