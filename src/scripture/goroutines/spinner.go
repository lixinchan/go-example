package goroutines

import (
	"fmt"
	"time"
)

// Spinner spinner
func Spinner() {
	go spinner(100 * time.Millisecond)
	x := 45
	fibResult := fib(x)
	fmt.Println("fib result:", fibResult)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
