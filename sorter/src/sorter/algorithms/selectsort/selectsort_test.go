// selectsort_test.go
package selectsort

func main() {
	var array []int = []int{5, 7, 6, 1, 4, 3, 2, 0}
	printArray(array)
	SelectSort(array)
	printArray(array)
}
