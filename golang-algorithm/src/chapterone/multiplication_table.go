package chapterone

import (
	"fmt"
	"strconv"
)

// multiplication table
func MultiplicationTable() {
	for idx := 1; idx <= 9; idx++ {
		for jdx := 1; jdx <= idx; jdx++ {
			var ret string
			if idx*jdx < 10 && jdx != 1 {
				ret = " " + strconv.Itoa(idx*jdx)
			} else {
				ret = strconv.Itoa(idx * jdx)
			}
			fmt.Print(jdx, " * ", idx, " = ", ret, " ")
		}
		fmt.Println()
	}
}
