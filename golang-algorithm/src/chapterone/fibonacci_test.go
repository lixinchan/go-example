package chapterone_test

import (
	. "chapterone"
	"testing"
	"fmt"
)

func TestFibonacci(t *testing.T) {
	for index := 0; index < 15; index++ {
		fmt.Println(index, ":", Fibonacci(index))
	}
}
