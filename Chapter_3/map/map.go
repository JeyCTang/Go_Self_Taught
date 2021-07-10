package main

import (
	"fmt"
	"sort"
)

func main() {
	scene := make(map[string]int)
	// prepare the data
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960

	// state a slice to save the data in map
	var sceneList []string
	// save the data in map to sceneList
	for k := range scene {
		sceneList = append(sceneList, k)
	}
	// sort the slice
	sort.Strings(sceneList)
	// output the sorted data
	fmt.Println(sceneList)

	for k, v := range scene {
		fmt.Println(k, v)
	}
	fmt.Println("-------------------")
	delete(scene, "brazil")
	for k, v := range scene {
		fmt.Println(k, v)
	}
}
