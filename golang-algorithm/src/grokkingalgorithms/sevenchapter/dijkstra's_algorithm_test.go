package sevenchapter

import (
	"testing"
)

func TestDijkstraAlgorithm(t *testing.T) {
	InitData()
	t.Log(DijkstraAlgorithm())
}
