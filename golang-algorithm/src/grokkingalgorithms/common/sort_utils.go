package common

import "fmt"

// compare v, w
func Less(v, w int) bool {
	return v-w < 0
}

// exchange array[idx], array[jdx]
func Exchange(array []int, idx, jdx int) {
	array[idx], array[jdx] = array[jdx], array[idx]
}

// print array
func PrintArray(array []int) {
	if array == nil {
		fmt.Println("nil")
		return
	}
	if len(array) == 0 {
		fmt.Println("[]")
		return
	}
	fmt.Println(array)
}
