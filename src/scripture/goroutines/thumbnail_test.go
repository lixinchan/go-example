package goroutines

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
	"testing"
)

func TestImageFile(t *testing.T) {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		thumb, err := ImageFile(input.Text())
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Println(thumb)
	}
	if err := input.Err(); err != nil {
		log.Fatal(err)
	}
}

func makeThumbnails(filenames []string) {
	for _, filename := range filenames {
		if _, err := ImageFile(filename); err != nil {
			log.Println(err)
		}
	}
}

// incorrect not wait finish, but return
func makeThumbnailsTwo(filenames []string) {
	for _, filename := range filenames {
		go ImageFile(filename)
	}
}

// use chan to wait finish
func makeThumbnailsThree(filenames []string) {
	finished := make(chan bool)
	for _, filename := range filenames {
		go func(filename string) {
			ImageFile(filename)
			finished <- true
		}(filename)
	}

	for range filenames {
		<-finished
	}
}

// return error incorrect
func makeThumbnailsFour(filenames []string) error {
	errors := make(chan error)
	for _, f := range filenames {
		go func(f string) {
			if _, err := ImageFile(f); err != nil {
				errors <- err
			}
		}(f)
	}

	for range filenames {
		if err := <-errors; err != nil {
			return err
		}
	}
	return nil
}

func makeThumbnailsFive(filenames []string) (thumbnails []string, err error) {
	type item struct {
		thumbnail string
		err       error
	}

	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbnail, it.err = ImageFile(f)
			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbnails = append(thumbnails, it.thumbnail)
	}
	return thumbnails, nil
}

func makeThumbnailsSix(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup
	for f := range filenames {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			thumb, err := ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(f)
	}

	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}
	return total
}
