package common

import (
	"bytes"
	"fmt"
)

// HashSet
type HashSet struct {
	m map[interface{}]bool
}

// init HashSet
func NewHashSet() *HashSet {
	return &HashSet{make(map[interface{}]bool)}
}

// add
func (hashSet *HashSet) Add(e interface{}) bool {
	if e == nil {
		return false
	}
	if hashSet.m[e] {
		return false
	}
	hashSet.m[e] = true
	return true
}

// add all
func (hashSet *HashSet) AddAll(array []interface{}) bool {
	if len(array) == 0 {
		return false
	}
	for _, v := range array {
		hashSet.Add(v)
	}
	return true
}

// remove element
func (hashSet *HashSet) Remove(e interface{}) {
	delete(hashSet.m, e)
}

// clear
func (hashSet *HashSet) Clear() {
	hashSet.m = make(map[interface{}]bool)
}

// is contains
func (hashSet *HashSet) Contains(e interface{}) bool {
	return hashSet.m[e]
}

// length
func (hashSet *HashSet) Length() int {
	return len(hashSet.m)
}

// same HashSet
func (hashSet *HashSet) Same(other *HashSet) bool {
	if other == nil {
		return false
	}
	if hashSet.Length() != other.Length() {
		return false
	}
	for key := range hashSet.m {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

// elements snapshot
func (hashSet *HashSet) Elements() []interface{} {
	initLen := hashSet.Length()
	snapshot := make([]interface{}, initLen)
	actualLen := 0
	for key := range hashSet.m {
		if initLen > actualLen {
			snapshot[actualLen] = key
		} else {
			snapshot = append(snapshot, key)
		}
		actualLen++
	}
	if actualLen < initLen {
		snapshot = snapshot[:actualLen]
	}
	return snapshot
}

// to String
func (hashSet *HashSet) ToString() string {
	var buffer bytes.Buffer
	buffer.WriteString("HashSet{")
	flag := 0
	for key := range hashSet.m {
		buffer.WriteString(fmt.Sprintf("%v", key))
		flag++
		if flag < hashSet.Length() {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("}")
	return buffer.String()
}

// is super set
func (hashSet *HashSet) IsSuperSet(other *HashSet) bool {
	if other == nil {
		return false
	}

	setLength := hashSet.Length()
	otherSetLength := other.Length()
	if setLength == 0 || setLength == otherSetLength {
		return false
	}
	if setLength > 0 && otherSetLength == 0 {
		return true
	}
	for _, v := range other.Elements() {
		if !hashSet.Contains(v) {
			return false
		}
	}
	return true
}

// set union
func (hashSet *HashSet) Union(other *HashSet) *HashSet {
	if hashSet == nil || other == nil {
		return nil
	}

	unionSet := NewHashSet()
	for _, v := range hashSet.Elements() {
		unionSet.Add(v)
	}
	if other.Length() == 0 {
		return unionSet
	}
	for _, v := range other.Elements() {
		unionSet.Add(v)
	}
	return unionSet
}

// set intersection
func (hashSet *HashSet) Intersection(other *HashSet) *HashSet {
	if hashSet == nil || other == nil {
		return nil
	}
	intersectionSet := NewHashSet()
	if hashSet.Length() == 0 || other.Length() == 0 {
		return intersectionSet
	}
	if hashSet.Length() > other.Length() {
		for _, v := range other.Elements() {
			if hashSet.Contains(v) {
				intersectionSet.Add(v)
			}
		}
	} else {
		for _, v := range hashSet.Elements() {
			if other.Contains(v) {
				intersectionSet.Add(v)
			}
		}
	}
	return intersectionSet
}

// Difference set
func (hashSet *HashSet) DifferenceSet(other *HashSet) *HashSet {
	if hashSet == nil || other == nil {
		return nil
	}
	differenceSet := NewHashSet()
	if other.Length() == 0 {
		for _, v := range hashSet.Elements() {
			differenceSet.Add(v)
		}
		return differenceSet
	}
	for _, v := range hashSet.Elements() {
		if !other.Contains(v) {
			differenceSet.Add(v)
		}
	}
	return differenceSet
}

// Symmetric difference set
func (hashSet *HashSet) SymmetricDiffSet(other *HashSet) *HashSet {
	if hashSet == nil || other == nil {
		return nil
	}

	diffSet := hashSet.DifferenceSet(other)
	if other.Length() == 0 {
		return diffSet
	}

	otherDiffSet := other.DifferenceSet(hashSet)
	return diffSet.Union(otherDiffSet)
}
