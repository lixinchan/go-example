package links

import (
	"os"
	"testing"
)

func TestBreadthFirst(t *testing.T) {
	BreadthFirst(Crawl, os.Args[1:])
}
