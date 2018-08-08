// basic_error
package basic

import (
	"fmt"
)

type error interface {
	Error() string
}

func Foo(param int) (n int, err error) {

}

type PathError struct {
	Op   string
	Path string
	Err  error
}

func (e *PathError) Error() string {
	return e.Op + " " + e.Path + ":" + e.Err.Error()
}

func main() {

	fmt.Println("Hello World!")
	//错误处理

	n := Foo(0)
	if err != nil {
		// 错误处理
	} else {
		// 正常流程
	}

	// 自定义error类型

}
