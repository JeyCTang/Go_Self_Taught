package main

import "fmt"

func main() {
	// initiate Player object, set the speed as 0.5
	p := NewPlayer(0.5)
	// move player to (3,1)
	p.MoveTo(Vec2{3, 1})
	// if the player isn't reach to target position, then loop the movement
	for !p.IsArrived() {
		// update the position of player
		p.Update()
		// print the position everytime after player moved
		fmt.Println(p.Pos())
	}
}
