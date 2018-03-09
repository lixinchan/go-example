// sortUtils
package main

import (
	"fmt"
)

func Less(v, w int) bool {
	return v-w < 0
}

func Exchage(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}

func PrintArray(array []int) {
	for _, v := range array {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func main() {
	fmt.Println("Hello World!")
}
