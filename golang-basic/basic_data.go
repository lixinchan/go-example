// basic_data.go
package main

import (
	"fmt"
	//	"os"
)

//func NewFile(fd int, name string) *File {
//	if fd < 0 {
//		return nil
//	}
//	f := new(File)
//	f.fd = fd
//	f.name = name
//	f.dirinfo = nil
//	f.nepipe = 0
//	return f
//}

// 表达式new File(),&File{}等价
//func NewFile(fd int, name string) *File {
//	if fd < 0 {
//		return nil
//	}
//	return &File{fd, name, nil, 0}
//}

//
//func NewFile(fd int, name string) *File {
//	if fd < 0 {
//		return nil
//	}
//	return &File{fd: name, name: name}
//}

func main() {

	//	slice := make([]int, 10, 100)
	//	fmt.Println("len:", len(slice))
	//	fmt.Println("cap:", cap(slice))
	// new 返回&,*array == nil
	//	var array *[]int = new([]int)
	//	fmt.Println(*array == nil)

	//	var v []int = make([]int, 100)
	//	fmt.Println("len:", len(v))
	//	fmt.Println("cap:", cap(v))

	// 以下写法等价
	//	var p *[]int = new([]int)
	//	*p = make([]int, 100, 100)

	//	p := make([]int, 100)

	// obtain pointer
	//	var array *[]int = new([]int)
	//	array := make([]int, 10)
	//	fmt.Println(&array)

}
