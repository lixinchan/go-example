package chapterone

// data type
type linkedStackData interface{}

// linked stack
type linkedStack struct {
	data linkedStackData
	next *linkedStack
}

// linkedStack's methods
type LinkedStack interface {
	Push(data linkedStackData)
	Pop() linkedStackData
	Size() int
	IsEmpty() bool
	IsFull() bool
}

// init linkedStack
func InitLinkedStack() *linkedStack {
	return &linkedStack{}
}

func (stack *linkedStack) Push(data linkedStackData) {

}

func (stack *linkedStack) Pop() linkedStackData {

	return nil
}

func (stack *linkedStack) Size() int {

	return 0
}

func (stack *linkedStack) IsEmpty() bool {

	return false
}

func (stack *linkedStack) IsFull() bool {

	return false
}
