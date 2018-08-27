package chapterone

import "errors"

type ArrayStack interface {
	Push(data interface{})     // push
	Pop() (interface{}, error) // pop
	Top() (interface{}, error) // top
	Size() int                 // size
	IsEmpty() bool             // empty
	IsFull() bool              // full
}

type StackArray []interface{}

// push
func (array *StackArray) Push(data interface{}) {
	*array = append(*array, data)
}

// pop
func (array *StackArray) Pop() (interface{}, error) {
	currentPointer := *array
	if len(currentPointer) == 0 {
		return nil, errors.New("index out of bound")
	}
	value := currentPointer[len(currentPointer)-1]
	*array = currentPointer[:len(currentPointer)-1]
	return value, nil
}

// top
func (array StackArray) Top() (interface{}, error) {
	if len(array) == 0 {
		return nil, errors.New("index out of bound")
	}
	return array[len(array)-1], nil
}

// size
func (array StackArray) Size() int {
	return len(array)
}

// isEmpty
func (array StackArray) IsEmpty() bool {
	return len(array) == 0
}

// isFull
func (array StackArray) IsFull() bool {
	return len(array) == cap(array)
}
