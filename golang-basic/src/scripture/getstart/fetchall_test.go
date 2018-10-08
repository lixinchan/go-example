package getstart

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestFetch(t *testing.T) {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go Fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
