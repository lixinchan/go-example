package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	var s []string
	for idx := range ip {
		s = append(s, strconv.Itoa(int(ip[idx])))
	}
	return strings.Join(s, ".")
}

func main() {
	hosts := map[string]IPAddr{
		"lo":          {127, 0, 0, 1},
		"dns":         {2, 2, 2, 2},
		"new-bee-dns": {1, 1, 1, 1},
	}
	for name, ip := range hosts {
		fmt.Printf("%v : %v\n", name, ip)
	}
}
