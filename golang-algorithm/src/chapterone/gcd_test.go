package chapterone_test

import (
	. "chapterone"
	"testing"
	"fmt"
)

func TestGcd(t *testing.T) {
	retVal := Gcd(16, 4)
	fmt.Println("The gcd is:", retVal)
}
