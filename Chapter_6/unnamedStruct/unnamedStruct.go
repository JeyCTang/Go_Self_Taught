package main

import "fmt"

// print the type of message, introduce unnamed struct
func printMsgType(msg *struct {
	id   int
	data string
}) {
	// print the type of msg
	fmt.Printf("%T\n", msg)
}
func main() {
	// initiate an unnamed struct
	msg := &struct {
		id   int
		data string
	}{1024, "hello"}

	printMsgType(msg)
}
