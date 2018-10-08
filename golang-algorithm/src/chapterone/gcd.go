// algorithm_gcd.go
package chapterone

// gcd recursive
func GcdRecursive(p, q int) int {
	if q == 0 {
		return p
	}
	r := p % q
	return GcdRecursive(q, r)
}

// gcd loop
func GcdLoop(p, q int) int {
	for q != 0 {
		p, q = q, p%q
	}
	return p
}
