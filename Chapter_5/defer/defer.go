package main

import "fmt"

func main() {
	fmt.Println("defer begin")

	// put defer into defer stack
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)

	fmt.Println("defer end")
}
