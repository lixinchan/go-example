package sevenchapter

import (
	"math"
	"strings"
)

var graph map[string]map[string]float64
var costs map[string]float64
var parents map[string]string
var processed []string

// init data
func InitData() {
	// init graph
	graph = make(map[string]map[string]float64)
	graph["start"] = map[string]float64{}
	graph["start"]["a"] = 6
	graph["start"]["b"] = 2
	graph["a"] = map[string]float64{}
	graph["a"]["finished"] = 1
	graph["b"] = map[string]float64{}
	graph["b"]["a"] = 3
	graph["b"]["finished"] = 5
	graph["finished"] = map[string]float64{}

	// init costs
	costs = make(map[string]float64)
	costs["a"] = 6
	costs["b"] = 2
	costs["finished"] = math.Inf(1)

	// init parents
	parents = make(map[string]string)
	parents["a"] = "start"
	parents["b"] = "start"
	parents["finished"] = " "
}

// main algorithm
func DijkstraAlgorithm() map[string]float64 {
	node := findLowestCostNode(costs)
	for node != " " {
		cost := costs[node]
		neighbors := graph[node]
		for key := range neighbors {
			newCost := cost + neighbors[key]
			if costs[key] > newCost {
				costs[key] = newCost
				parents[key] = node
			}
		}
		processed = append(processed, node)
		node = findLowestCostNode(costs)
	}
	return costs
}

// find lowest cost node
func findLowestCostNode(costs map[string]float64) string {
	lowestCost := math.Inf(1)
	lowestCostNode := " "
	for node := range costs {
		cost := costs[node]
		if (cost < lowestCost) && !findProcessed(processed, node) {
			lowestCost = cost
			lowestCostNode = node
		}
	}
	return lowestCostNode
}

// find node processed
func findProcessed(processed []string, node string) bool {
	for idx := range processed {
		if strings.Compare(processed[idx], node) == 0 {
			return true
		}
	}
	return false
}
