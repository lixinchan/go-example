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
