package eightchapter

import (
	"fmt"
	"grokkingalgorithms/common"
)

var (
	//all station
	stations map[string]*common.HashSet
	// states needed
	statesNeeded *common.HashSet
	// final station
	finalStations *common.HashSet
)

// init data
func InitData() {
	// init stations
	stations = make(map[string]*common.HashSet)
	keyOne := make([]interface{}, 0, 10)
	keyOne = append(keyOne, "id", "nv", "ut")
	oneSet := common.NewHashSet()
	oneSet.AddAll(keyOne)
	stations["keyOne"] = oneSet
	keyTwo := make([]interface{}, 0, 10)
	keyTwo = append(keyTwo, "wa", "id", "mt")
	twoSet := common.NewHashSet()
	twoSet.AddAll(keyTwo)
	stations["keyTwo"] = twoSet
	keyThree := make([]interface{}, 0, 10)
	keyThree = append(keyThree, "or", "nv", "ca")
	threeSet := common.NewHashSet()
	threeSet.AddAll(keyThree)
	stations["keyThree"] = threeSet
	keyFour := make([]interface{}, 0, 10)
	keyFour = append(keyFour, "nv", "ut")
	fourSet := common.NewHashSet()
	fourSet.AddAll(keyFour)
	stations["keyFour"] = fourSet
	keyFive := make([]interface{}, 0, 10)
	keyFive = append(keyFive, "ca", "az")
	fiveSet := common.NewHashSet()
	fiveSet.AddAll(keyFive)
	stations["keyFive"] = fiveSet

	// init states needed
	neededStates := make([]interface{}, 0, 10)
	neededStates = append(neededStates, "mt", "wa", "or", "id", "nv", "ut", "ca", "az")
	statesNeeded = common.NewHashSet()
	statesNeeded.AddAll(neededStates)
}

// main algorithm
func GreedyAlgorithm() {
	finalStations = common.NewHashSet()
	for statesNeeded.Length() > 0 {
		var bestStation string
		statesCovered := common.NewHashSet()
		for station, states := range stations {
			covered := statesNeeded.DifferenceSet(states)
			if covered.Length() > statesCovered.Length() {
				bestStation = station
				statesCovered = covered
			}
			statesNeeded = statesNeeded.Intersection(statesCovered)
			if len(bestStation) > 0 {
				finalStations.Add(bestStation)
			}
		}
	}
	fmt.Println(finalStations)
}
