package main

import "fmt"

func main() {
	// prepare a string type
	var house = "Malibu Point 10880, 90265"
	// get the add from the above string, the type of ptr is *string
	ptr := &house
	// print the type of ptr
	fmt.Printf("ptr type: %T\n", ptr)
	// print the address of ptr
	fmt.Printf("address: %p\n", ptr)
	// get the value of pointer
	value := *ptr
	// the Type of variable
	fmt.Printf("value type: %T\n", value)
	// after we use *ptr, what we get is the value of variable
	fmt.Printf("value: %s\n", value)
}
