package sixthchapter

import (
	"fmt"
	"strings"
)

var graph map[string][]string

// init graph
func InitGraph() {
	graph = make(map[string][]string)
	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}
}

// mango seller
func isMangoSeller(name string) bool {
	return strings.HasSuffix(name, "m")
}

// contains searched
func containsSearched(name string, searched []string) bool {
	for _, ele := range searched {
		if strings.EqualFold(name, ele) {
			return true
		}
	}
	return false
}

// breadth-first search
func Search(name string) bool {
	if strings.TrimSpace(name) == "" {
		return false
	}

	var searchQueue []string
	searchQueue = append(searchQueue, graph[name] ...)

	// searched vertice
	searched := make([]string, 10)

	var person string
	for len(searchQueue) != 0 {
		person = searchQueue[0]
		searchQueue = searchQueue[1:]
		if !containsSearched(person, searched) {
			if isMangoSeller(person) {
				fmt.Println(person, "is mango seller.")
				return true
			} else {
				searchQueue = append(searchQueue, graph[person] ...)
				searched = append(searched, person)
			}
		}
	}
	return false
}
