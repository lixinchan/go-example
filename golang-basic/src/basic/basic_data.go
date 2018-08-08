// basic_data.go
package basic

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

// array
func Sum(array *[3]float64) (sum float64) {
	for _, v := range array {
		sum += v
	}
	return
}

// slice
//func (f *File) Read(buf []byte) (n int, err error) {
//	// 读取前32个字节
//	//	n, err := f.Read(buf[0:32])

//	for i := 0; i < 32; i++ {
//		nbytes, e := f.Read(buf[i : i+1])
//		if nbytes == 0 || e != nil {
//			err = e
//			break
//		}
//		n += nbytes
//	}
//	return
//}

// test append
func myAppend(slice, data []byte) []byte {
	length := len(slice)
	if length+len(data) > cap(slice) {
		newSlice := make([]byte, (length+len(data))*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : length+len(data)]
	for i, v := range data {
		slice[length+i] = v
	}
	return slice
}

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

	//array
	//	array := [...]float64{6.0, 7.1, 8.2}
	//	fmt.Println(Sum(&array))

	//self define slice append
	//	slice1 := []byte{1, 2, 3}
	//	slice2 := []byte{4, 5, 6}
	//	slice3 := myAppend(slice1, slice2)
	//	fmt.Println(slice3)

	// two-dimensional array
	// array-of-arrays
	type Transform [8][8]float64
	// slice-of-slices
	type LineOfText [][]byte

	var XSize int = 3
	var YSize int = 4
	//	picture := make([][]uint8, YSize)
	//	for i := range picture {
	//		picture[i] = make([]uint8, XSize)
	//	}

	picture := make([][]uint8, YSize)
	pixels := make([]uint8, XSize*YSize)
	for i := range picture {
		picture[i], pixels = pixels[:XSize], pixels[XSize:]
	}
}
