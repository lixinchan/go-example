package common

import "testing"

func TestHashSet_Add(t *testing.T) {
	hashSet := NewHashSet()
	for idx := 0; idx < 10; idx++ {
		hashSet.Add(idx)
	}
	t.Log(hashSet.ToString())
	t.Log(hashSet.Elements())
}
