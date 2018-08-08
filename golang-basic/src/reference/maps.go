package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

// equal two map
func equal(x, y map[int]string) bool {
	if len(x) != len(y) {
		return false
	}
	for key, xVal := range x {
		if yVal, ok := y[key]; !ok || xVal != yVal {
			return false
		}
	}
	return true
}

func main() {
	m = make(map[string]Vertex)
	m["test-map"] = Vertex{
		3213.12313, -234.3424,
	}
	fmt.Println(m["test-map"])
}
