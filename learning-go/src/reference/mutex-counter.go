package main

import (
	"sync"
	"fmt"
	"time"
)

type safeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// increment
func (counter *safeCounter) increment(key string) {
	counter.mux.Lock()
	counter.v[key]++
	counter.mux.Unlock()
}

// value
func (counter *safeCounter) value(key string) int {
	counter.mux.Lock()
	defer counter.mux.Unlock()
	return counter.v[key]
}

func main() {
	counter := safeCounter{v: make(map[string]int)}
	for idx := 0; idx < 100000000; idx++ {
		go counter.increment("testKey")
	}
	time.Sleep(500 * time.Millisecond)
	fmt.Println(counter.value("testKey"))
}
