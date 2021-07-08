package main

import "fmt"

// NewInt define NewInt as int type
type NewInt int

// IntAlias The other name of int is IntAlias
type IntAlias = int

func main() {
	// state a is a NewInt type
	var a NewInt
	// check the type of a
	fmt.Printf("a type: %T\n", a)

	// state a2 as IntAlias type
	var a2 IntAlias
	// check the type of a2
	fmt.Printf("a2 type: %T\n", a2)
}
