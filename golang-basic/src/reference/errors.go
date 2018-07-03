package main

import (
	"strconv"
	"fmt"
	"time"
)

func testErr(str string) {
	idx, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("can not convert number: %v\n", err)
		return
	}
	fmt.Printf("convert integer: %v\n", idx)
}

type myError struct {
	when time.Time
	what string
}

func (e *myError) Error() string {
	return fmt.Sprintf("at %v, %s", e.when, e.what)
}

func run() error {
	return &myError{
		time.Now(),
		"It didn't work.",
	}
}

func main() {
	testErr("21")

	if err := run(); err != nil {
		fmt.Println(err)
	}
}
