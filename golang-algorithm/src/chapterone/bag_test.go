package chapterone_test

import (
	"testing"
	. "chapterone"
)

func TestAdd(t *testing.T) {
	bag := InitBag()
	if !bag.IsEmpty() {
		for idx := 10; idx > 0; idx-- {
			bag.Add(idx)
		}
	}
	t.Log(bag.Size())
}
