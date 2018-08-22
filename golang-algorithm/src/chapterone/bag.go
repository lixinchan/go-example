package chapterone

// data type
type bagData interface{}

// Bag's methods
type Bag interface {
	Size() int
	IsEmpty() bool
	Add(data bagData)
}

// type bagNode
type bagNode struct {
	data bagData
	next *bagNode
}

var (
	length int
)

// init bag
func InitBag() *bagNode {
	return &bagNode{}
}

// add
func (head *bagNode) Add(data bagData) {
	oldFirst := head
	newHead := bagNode{data: data}
	newHead.next = oldFirst
	length++
}

// size
func (head *bagNode) Size() int {
	return length
}

// isEmpty
func (head *bagNode) IsEmpty() bool {
	first := head
	return first == nil
}
