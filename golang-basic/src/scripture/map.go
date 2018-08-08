package scripture

import (
	"sort"
	"fmt"
)

// sort map by key
func SortMap(ages map[string]int) {
	if ages == nil {
		return
	}
	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}
