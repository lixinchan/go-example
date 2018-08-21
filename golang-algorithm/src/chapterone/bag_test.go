package chapterone_test

import (
	"testing"
	. "chapterone"
)

func TestAdd(t *testing.T) {
	t.Log(IsEmpty())
	for idx:=10; idx >0;idx-- {
		Add(idx)
	}
	t.Log(Size())
}
