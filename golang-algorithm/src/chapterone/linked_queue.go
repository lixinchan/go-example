// FIFO QUEUE
package chapterone

import (
	"errors"
	"sync"
)

// node
type LinkedQueueNode struct {
	next *LinkedQueueNode // next
	data interface{}      // data
	lock sync.RWMutex     // lock
}

// interface
type LinkedQueue interface {
	LinkedEnqueue(data interface{})      // enqueue
	LinkedDequeue() (interface{}, error) // dequeue
	LinkedIsEmpty() bool                 // isEmpty
	LinkedSize() int                     // size
	//LinkedClear() bool                   // clear
}

var (
	linkedSize   int              // size
	linkedHeader *LinkedQueueNode // header
	linkedTail   *LinkedQueueNode // tail
)

// init linked queue
func InitLinkedQueue() *LinkedQueueNode {
	node := &LinkedQueueNode{next: nil, data: nil}
	linkedHeader = node
	linkedTail = node
	return node
}

// enqueue
func (node *LinkedQueueNode) LinkedEnqueue(data interface{}) {
	node.lock.Lock()
	if node.LinkedIsEmpty() {
		node.data = data
		node.next = nil
	} else {
		oldTail := linkedTail
		linkedTail = &LinkedQueueNode{next: nil, data: data}
		oldTail.next = linkedTail
	}
	linkedSize++
	node.lock.Unlock()
}

// dequeue
func (node *LinkedQueueNode) LinkedDequeue() (interface{}, error) {
	node.lock.Lock()
	if node.LinkedIsEmpty() {
		node.lock.Unlock()
		return nil, errors.New("queue is nil")
	}
	oldHeader := linkedHeader
	linkedHeader = linkedHeader.next
	linkedSize--
	node.lock.Unlock()
	return oldHeader.data, nil
}

// isEmpty
func (node LinkedQueueNode) LinkedIsEmpty() bool {
	return linkedSize == 0
}

// size
func (node LinkedQueueNode) LinkedSize() int {
	return linkedSize
}
