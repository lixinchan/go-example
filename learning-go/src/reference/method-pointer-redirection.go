package main

import "fmt"

type matrixForRedirect struct {
	x, y float64
}

func (m *matrixForRedirect) scale(f float64) {
	m.x = m.x * f
	m.y = m.y * f
}

func scaleFunc(m *matrixForRedirect, f float64) {
	m.x = m.x * f
	m.y = m.y * f
}

func main() {
	v := matrixForRedirect{3, 4}
	(&v).scale(10)
	v.scale(1)
	scaleFunc(&v, 10)

	p := &matrixForRedirect{5, 12}
	p.scale(10)
	scaleFunc(p, 10)
	fmt.Println(v, p)
}
