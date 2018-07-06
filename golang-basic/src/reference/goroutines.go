package main

import (
	"time"
	"fmt"
)

func say(s string) {
	for idx := 0; idx < 500; idx++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
