package main

import (
	"os"
	"io"
	"fmt"
)

func copyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

func deferTest() (i int) {
	defer func() { i++ }()
	return 1
}

func main() {
	//copyFile("", "")
	fmt.Println(deferTest())
}
