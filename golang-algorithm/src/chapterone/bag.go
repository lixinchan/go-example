package chapterone

type Bag interface {
	Size() int
	IsEmpty() bool
	Add()
}

// type node
type node struct {
	data interface{}
	next *node
}

var (
	first  node
	length int
)

// add
func Add(data interface{}) {
	oldFirst := first
	first = node{}
	first.data = data
	first.next = &oldFirst
	length++
}

// size
func Size() int {
	return length
}

// isEmpty
func IsEmpty() bool {
	return length == 0
}
