// FIFO QUEUE
package chapterone

type LinkedQueueNode struct {
	next *LinkedStackNode
	data interface{}
}

type LinkedQueue interface {
	LinkedEnqueue(data interface{})
	LinkedDequeue() (interface{}, error)
	LinkedIsEmpty() bool
	LinkedSize() int
}

var (
	linkedSize   int
	linkedHeader *LinkedQueueNode
	linkedTail   *LinkedQueueNode
)

func (node *LinkedQueueNode) LinkedEnqueue(data interface{}) {

}

func (node *LinkedQueueNode) LinkedDequeue() (interface{}, error) {

	return nil, nil
}

func (node LinkedQueueNode) LinkedIsEmpty() bool {
	return linkedSize == 0
}

func (node LinkedQueueNode) LinkedSize() int {
	return linkedSize
}
