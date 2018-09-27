package chapterone

import (
	"errors"
	"sync"
)

type ArrayStack interface {
	Push(data interface{})     // push
	Pop() (interface{}, error) // pop
	Top() (interface{}, error) // top
	Size() int                 // size
	IsEmpty() bool             // empty
	IsFull() bool              // full
}

type StackArray struct {
	element [] interface{}
	lock    sync.RWMutex
}

// push
func (array *StackArray) Push(data interface{}) {
	array.lock.Lock()
	array.element = append(array.element, data)
	array.lock.Unlock()
}

// pop
func (array *StackArray) Pop() (interface{}, error) {
	array.lock.RLock()
	currentPointer := array.element
	if len(currentPointer) == 0 {
		return nil, errors.New("index out of bound")
	}
	value := currentPointer[len(currentPointer)-1]
	array.element = currentPointer[:len(currentPointer)-1]
	array.lock.Unlock()
	return value, nil
}

// top
func (array StackArray) Top() (interface{}, error) {
	array.lock.RLock()
	if len(array.element) == 0 {
		return nil, errors.New("index out of bound")
	}
	array.lock.RUnlock()
	return array.element[len(array.element)-1], nil
}

// size
func (array StackArray) Size() int {
	return len(array.element)
}

// isEmpty
func (array StackArray) IsEmpty() bool {
	return len(array.element) == 0
}

// isFull
func (array StackArray) IsFull() bool {
	return len(array.element) == cap(array.element)
}
