// basic_map
package basic

import (
	"fmt"
)

//type PersonInf struct {
//	ID      string
//	Name    string
//	Address string
//}

func Maps() {
	//	var personDB map[string]PersonInf
	//	personDB = make(map[string]PersonInf)
	//	personDB["one"] = PersonInf{"1", "Jack", "Earth"}
	//	personDB["two"] = PersonInf{"2", "Rose", "Earth"}
	//	person, ok := personDB["three"]
	//	if ok {
	//		fmt.Println("The element found:", person)
	//	} else {
	//		fmt.Println("Not found")
	//	}

	// delete() func
	//	delete(personDB, "one")
	//	fmt.Println(personDB)

	// element find
	//	person1, ok := personDB["two"]
	//	fmt.Println(person1)

	var timeZone = map[string]int{"UTC": 0 * 60 * 60, "CST": -6 * 60 * 60}
	offset := timeZone["CST..d"]

	fmt.Println(offset)

	//	if off, ok := timeZone["UTC"]; ok {
	//		fmt.Println(ok)
	//		fmt.Println(off)
	//	}
}
