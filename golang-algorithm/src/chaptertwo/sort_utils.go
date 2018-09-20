// sort_utils
package chaptertwo

import (
	"fmt"
)

// compare element
func Less(v, w int) bool {
	return v-w < 0
}

// exchange element
func Exchange(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}

// print array
func PrintArray(array []int) {
	if array == nil {
		fmt.Print("nil")
		return
	}
	if len(array) == 0 {
		fmt.Print("[]")
		return
	}
	for _, v := range array {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
