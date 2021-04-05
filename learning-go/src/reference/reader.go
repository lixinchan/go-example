package main

import (
	"strings"
	"fmt"
	"io"
)

func main() {
	reader := strings.NewReader("Hello,Reader!")
	bytes := make([]byte, 15)
	for {
		n, err := reader.Read(bytes)
		fmt.Printf("n=%v, err=%v, bytes=%v\n", n, err, bytes)
		fmt.Printf("bytes[:n]=%q\n", bytes[:n])
		if err == io.EOF {
			break
		}
	}
}
