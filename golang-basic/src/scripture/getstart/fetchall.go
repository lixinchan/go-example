package getstart

import (
	"time"
	"net/http"
	"fmt"
	"io"
	"io/ioutil"
)

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	seconds := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", seconds, nbytes, url)
}
