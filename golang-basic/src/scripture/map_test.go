package scripture_test

import (
	. "scripture"
	"testing"
)

func TestSortMap(t *testing.T) {
	ages := map[string]int{
		"lmm": 28,
		"abb": 1,
		"clx": 27,
	}
	SortMap(ages)
}

func TestEqualMap(t *testing.T) {
	x := map[string]int{
		"lmm": 28,
		"abb": 1,
		"clx": 27,
	}
	y := map[string]int{
		"lmm": 28,
		"abb": 10,
		"clx": 27,
	}
	t.Log(EqualMap(x, y))
}
