package chapterone

type arrayStackData interface{}

type Stack interface {
	Push(data arrayStackData) // push bagData
	Pop() arrayStackData      // pop bagData
	Size() int                // size
	IsEmpty() bool            // empty
	IsFull() bool             // full
}

type arrayStack []arrayStackData

// init array stack
func InitArrayStack(size int) [size]arrayStackData {
	return [size]arrayStackData{}
}

// push
func (array *arrayStack) Push(data arrayStackData) {

}

// pop
func (array *arrayStack) Pop() arrayStackData {

}

// size
func (array *arrayStack) Size() int {
	return 0
}

// isEmpty
func (array *arrayStack) IsEmpty() bool {

	return false
}

// isFull
func (array *arrayStack) IsFull() bool {

	return false
}

// resizing
func resizing() []arrayStack {

	return nil
}
