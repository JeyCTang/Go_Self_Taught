package main

import "fmt"

func main() {
	defer fmt.Println("The 1st thing after panic")
	defer fmt.Println("The 2nd thing after panic")
	panic("panic!!!!")
}
