package main

import (
	"fmt"
	"sync"
)

func main() {
	var scene sync.Map
	// save key and value to sync.Map
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)
	// get the value according to the key
	fmt.Println(scene.Load("london"))
	value, result := scene.Load("london")
	fmt.Println(value, result)
	// delete the key&value pairs according to the key
	scene.Delete("london")
	// print all the current key&value
	scene.Range(func(key, value interface{}) bool {
		fmt.Println("iterate:", key, value)
		return true
	})
}
