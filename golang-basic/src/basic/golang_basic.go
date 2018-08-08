package basic

import (
	"fmt"
)

//func modify(arr [5]int) {
//	arr[0] = 10
//	fmt.Println("In modify array values:", arr)
//}

//func reverseArray(array []int) {
//	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
//		array[i], array[j] = array[j], array[i]
//	}
//}

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	//	var str string
	//	str := "a b"
	//	ch := str[0]
	//	fmt.Printf("%s\n", str)
	//	fmt.Printf("%c\n", ch)
	//	fmt.Printf("len=%d\n", len(str))

	//	len := len(str)

	//	for i := 0; i < len; i++ {
	//		ch := str[i]
	//		fmt.Println(i, ch)
	//	}

	//	for i, chr := range str {
	//		fmt.Println(i, chr)
	//	}

	var arr []int = []int{1, 2, 3, 4, 5}

	//	for i := 0; i < len(arr); i++ {
	//		fmt.Println("Element", i, "of array is", arr[i])
	//	}

	//	for i, v := range arr {
	//		fmt.Println("Element", i, "of array is", v)
	//	}

	//	modify(arr)
	//	fmt.Println("In main array values:", arr)

	//	 slice based on array
	//	var mySlice []int = arr[1:3]
	mySlice := arr[1:3]
	fmt.Println(mySlice)

	//	for _, v := range mySlice {
	//		fmt.Println("mySlice:", v)
	//	}

	// slice based on new
	//	mySlice := make([]int, 5)
	//	mySlice := make([]int, 5, 10)
	//	mySlice = []int{1, 2, 3, 4, 5}
	//	mySlice := []int{1, 2, 3, 4, 5}
	//	for i, v := range mySlice {
	//		fmt.Println("mySlice[", i, "]", v)
	//	}
	//	fmt.Println("mySlice's length:", len(mySlice))
	//	fmt.Println("mySlice's capacity:", cap(mySlice))

	// reverse array
	//	reverseArray(arr)
	//	fmt.Println(arr)

	// append array
	//	mySlice = append(mySlice, 5, 4, 3, 2, 1)
	//	fmt.Println("mySlice:", mySlice)

	//	mySlice2 := arr[:3]
	//	mySlice2 = append(mySlice2, mySlice...)
	//	fmt.Println(mySlice2)

	//	fmt.Println(mySlice)
	//	mySlice3 := mySlice[:10]
	//	fmt.Println("mySlice3:", mySlice3)

	//	slice1 := []int{1, 2, 3, 4, 5}
	//	slice2 := []int{5, 4, 3}
	//	copy(slice2, slice1)
	//	fmt.Println(slice1)
	//	fmt.Println(slice2)

	//	var persionDB map[string]PersonInfo
	//	persionDB = make(map[string]PersonInfo)

	//	persionDB["1"] = PersonInfo{"1", "Jack", "earth..."}
	//	persionDB["2"] = PersonInfo{"2", "Rose", "earth..."}

	//	delete(persionDB, "1")

	//	person, ok := persionDB["1"]
	//	if ok {
	//		fmt.Println(person.ID, person.Name, person.Address)
	//	} else {
	//		fmt.Println("person not found~")
	//	}
}
