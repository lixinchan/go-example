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

func TestHashSet_AddAll(t *testing.T) {
	hashSet := NewHashSet()
	array := make([]interface{}, 0, 10)
	for idx := 0; idx < 10; idx++ {
		array = append(array, idx)
	}
	hashSet.AddAll(array)
	t.Log(hashSet.ToString())
}
