package main

import (
	"fmt"
	"reflect"
)

type Player struct {
	Name        string
	HealthPoint int
	MagicPoint  int
}
type Point struct {
	X int
	Y int
}

func main() {
	var p Point
	p.X = 20
	p.Y = 10
	fmt.Println(reflect.TypeOf(p))

	tank := new(Player)
	tank.Name = "Cannon"
	tank.HealthPoint = 300
	fmt.Println(reflect.TypeOf(tank))
}
