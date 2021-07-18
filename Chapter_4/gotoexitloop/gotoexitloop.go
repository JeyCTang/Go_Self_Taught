package main

import "fmt"

func main() {
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if y == 2 {
				// go to the label
				goto breakHere
			}
		}
	}
	// manually return, so the program won't execute breakHere sequentially
	return
	// label
breakHere:
	fmt.Println("Done")
}
