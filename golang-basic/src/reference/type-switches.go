package main

import "fmt"

func doSwitch(inter interface{}) {
	switch val := inter.(type) {
	case int:
		fmt.Printf("triple %v is %v\n", val, val*3)
	case string:
		fmt.Printf("%q is %d bytes long\n", val, len(val))
	default:
		fmt.Println("unknown type.", val)
	}
}

func main() {
	doSwitch(12)
	doSwitch("hello")
	doSwitch(false)
}
