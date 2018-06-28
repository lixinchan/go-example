package main

import "fmt"

type VertexLiterals struct {
	Lat, Long float64
}

var mapLiterals = map[string]VertexLiterals{
	"test-map": {3213.12313, -234.3424},
	"test":     {3213.12313, -234.3424},
}

func main() {
	fmt.Println(mapLiterals)
}
