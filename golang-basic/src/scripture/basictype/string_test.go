package basictype

import "testing"

func TestHasPrefix(t *testing.T) {
	s := "hello, world"
	substr := "hello"
	t.Log(HasPrefix(s, substr))
}

func TestHasSuffix(t *testing.T) {
	s := "hello, world"
	substr := "world"
	t.Log(HasSuffix(s, substr))
}

func TestContains(t *testing.T) {
	s := "hello, world"
	substr := "llo"
	t.Log(Contains(s, substr))
}
