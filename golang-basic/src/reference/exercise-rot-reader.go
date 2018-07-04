package main

import (
	"io"
	"strings"
	"os"
)

type rot13Reader struct {
	reader io.Reader
}

func (reader *rot13Reader) Read(b []byte) (n int, err error) {

}

func main() {
	str := strings.NewReader("clx")
	reader := rot13Reader{str}
	io.Copy(os.Stdout, &reader)
}
