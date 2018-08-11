package basic

import (
	"fmt"
)

type testError interface {
	Error() string
}

func Foo(param int) (n int, err testError) {
	if param <= 0 {
		return 0, err
	}
	return param, nil
}

type PathError struct {
	Op   string
	Path string
	Err  error
}

func (e *PathError) Error() string {
	return e.Op + " " + e.Path + ":" + e.Err.Error()
}

func Error() {

	fmt.Println("Hello World!")
	//错误处理

	n, err := Foo(0)
	if err != nil {
		// 错误处理
	} else {
		// 正常流程
		fmt.Print(n)
	}

	// 自定义error类型

}
