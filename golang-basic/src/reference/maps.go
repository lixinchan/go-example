package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["test-map"] = Vertex{
		3213.12313, -234.3424,
	}
	fmt.Println(m["test-map"])
}
