package main

import "fmt"

func main() {
	// columns
	for y := 1; y <= 9; y++ {
		// rows
		for x := 1; x <= y; x++ {
			fmt.Printf("%d * %d = %d  ", x, y, x*y)
		}
		// line break
		fmt.Println()
	}
}
