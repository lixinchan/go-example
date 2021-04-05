package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on:")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("MacOS")
	case "linux":
		fmt.Println("Linux")
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Printf("%s", os)
	}
}
