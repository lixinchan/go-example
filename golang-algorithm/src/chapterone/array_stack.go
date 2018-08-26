package chapterone

import "errors"

type Stack interface {
	Push(data interface{})     // push
	Pop() (interface{}, error) // pop
	Top() (interface{}, error) // top
	Size() int                 // size
	IsEmpty() bool             // empty
	IsFull() bool              // full
}

type ArrayStack []interface{}

// push
func (array *ArrayStack) Push(data interface{}) {
	*array = append(*array, data)
}

// pop
func (array *ArrayStack) Pop() (interface{}, error) {
	currentPointer := *array
	if len(currentPointer) == 0 {
		return nil, errors.New("index out of bound")
	}
	value := currentPointer[len(currentPointer)-1]
	*array = currentPointer[:len(currentPointer)-1]
	return value, nil
}

// top
func (array ArrayStack) Top() (interface{}, error) {
	if len(array) == 0 {
		return nil, errors.New("index out of bound")
	}
	return array[len(array)-1], nil
}

// size
func (array ArrayStack) Size() int {
	return len(array)
}

// isEmpty
func (array ArrayStack) IsEmpty() bool {
	return len(array) == 0
}

// isFull
func (array ArrayStack) IsFull() bool {
	return len(array) == cap(array)
}
