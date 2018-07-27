// sort_insertion
package chaptertwo

func Sort(array []int) {
	length := len(array)
	for i := 0; i < length; i++ {
		for j := i; j > 0 && Less(array[j], array[j-1]); j-- {
			Exchange(array, j, j-1)
		}
	}

}

func SortEnhance(array []int) {
	length := len(array)
	var j int
	for i := 1; i < length; i++ {
		tmp := array[i]
		for j = i - 1; j >= 0 && Less(tmp, array[j]); j-- {
			array[j+1] = array[j]
		}
		array[j+1] = tmp
	}
}
