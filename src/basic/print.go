package basic

import (
	"fmt"
	"os"
)

type T struct {
	a int
	b float64
	c string
}

// 控制自定义类型的默认格式
func (t *T) String() string {
	return fmt.Sprintf("%d%g%q", t.a, t.b, t.c)
}

func Print() {

	fmt.Printf("Hello %d\n", 23)
	fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
	fmt.Println("Hello", 23)
	fmt.Println(fmt.Sprint("Hello ", 23))

	fmt.Println("----------")

	var x uint64 = 1<<64 - 1
	fmt.Printf("%d %x; %d %x\n", x, x, int64(x), int64(x))

	fmt.Println("----------")

	var timeZone = map[string]int{"UTC": 0 * 60 * 60, "CST": -6 * 60 * 60}

	t := &T{7, -2.35, "abc\tdef"}
	// pring value
	fmt.Printf("%v\n", t)
	// 为结构体添加字段名称
	fmt.Printf("%+v\n", t)
	// 按照go的语法打印值
	fmt.Printf("%#v\n", t)
	fmt.Printf("%#v\n", timeZone)

	fmt.Println("----------")

	//print type
	fmt.Printf("%T\n", timeZone)

	fmt.Println("----------")

	// 遇到string/[]byte值时，产生带引号的字符串
	str := "hello, go"
	array := []byte{1, 2, 3, 4, 5}
	fmt.Printf("%q\n", str)
	fmt.Printf("%#q\n", array)

	fmt.Println("----------")

	// 控制自定义类型的默认格式

	fmt.Println(^0)
	fmt.Println(uint(0))
	fmt.Println(^uint(0))
	fmt.Println(^uint(0) >> 1)
}
