package chapterone

// linked stack
type LinkedStackNode struct {
	data interface{}
	next *LinkedStackNode
}

// linkedStack's methods
type LinkedStack interface {
	Push(data interface{})     // push
	Pop() (interface{}, error) // pop
	Top() (interface{}, error) // top
	Size() int                 // size
	IsEmpty() bool             // isEmpty
	IsFull() bool              // isFull
}

// push
func (stack *LinkedStackNode) Push(data interface{}) {

}

// pop
func (stack *LinkedStackNode) Pop() (interface{}, error) {

	return nil, nil
}

// top
func (stack *LinkedStackNode) Top() (interface{}, error) {

	return nil, nil
}

// size
func (stack LinkedStackNode) Size() int {

	return 0
}

// isEmpty
func (stack LinkedStackNode) IsEmpty() bool {

	return false
}

// isFull
func (stack LinkedStackNode) IsFull() bool {

	return false
}
