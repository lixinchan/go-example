// uses a map to record the number of times Add has been called with a given
// list of strings.
package compositetypes

import "fmt"

var localMap = make(map[string]int)

// add
func Add(list []string) {
	localMap[convert(list)]++
}

// convert slice to string
func convert(list []string) string {
	return fmt.Sprintf("%q", list)
}

// count
func Count(list []string) int {
	return localMap[convert(list)]
}
