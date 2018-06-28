package main

import "fmt"

type VertexLiteral struct {
	Lat, Long float64
}

var mapLiteral = map[string]VertexLiteral{
	"test-map": VertexLiteral{
		3213.12313, -234.3424,
	},
	"test": VertexLiteral{
		3213.12313, -234.3424,
	},
}

func main() {
	fmt.Println(mapLiteral)
}
