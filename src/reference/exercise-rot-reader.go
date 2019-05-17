package main

import (
	"io"
	"strconv"
	"fmt"
)

type rot13Reader struct {
	reader io.Reader
}

// init rot13
func initROT13() (rotMap map[string]string) {
	rotMap = make(map[string]string)
	var idx int
	var count int = 0
	for idx = 'a'; idx <= 'z'; idx++ {
		if count < 13 {
			rotMap[strconv.Itoa(idx)] = strconv.Itoa(idx + 13)
		} else {
			rotMap[strconv.Itoa(idx)] = strconv.Itoa(idx - 13)
		}
		count++
	}
	count = 0
	for idx = 'A'; idx <= 'Z'; idx++ {
		if count < 13 {
			rotMap[strconv.Itoa(idx)] = strconv.Itoa(idx + 13)
		} else {
			rotMap[strconv.Itoa(idx)] = strconv.Itoa(idx - 13)
		}
		count++
	}
	return rotMap
}

func (reader *rot13Reader) Read(b []byte) (n int, err error) {

	return 0, nil
}

func main() {
	//str := strings.NewReader("clx")
	//reader := rot13Reader{str}
	//io.Copy(os.Stdout, &reader)

	rotMap := initROT13()
	for key, val := range rotMap {
		fmt.Println(key, val)
	}
}
