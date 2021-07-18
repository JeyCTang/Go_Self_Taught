package main

import (
	"fmt"
)

// create a playerGen, input the player name, output a generator
func playerGen(name string) func() (string, int) {
	// hp is always 150
	hp := 150
	// return closure
	return func() (string, int) {
		// refer the variable into closure
		return name, hp
	}
}

func main() {
	// create a player generator
	generator := playerGen("high_noon")
	// return the hp and player name
	name, hp := generator()
	// print it out
	fmt.Println(name, hp)
}
