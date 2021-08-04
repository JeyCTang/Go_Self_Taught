package main

import (
	"fmt"
	"reflect"
)

// Point define Point struct
type Point struct {
	X int
	Y int
}

// Add un-pointer add method
func (p Point) Add(other Point) Point {
	// member values add together and return a new struct
	return Point{p.X + other.X, p.Y + other.Y}
}

func main() {
	// initiate point
	p1 := Point{1, 1}
	p2 := Point{2, 2}

	// one point plus with the other point
	result := p1.Add(p2)

	// print the result
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}
