package main

import "fmt"

// Invoker Interface of caller
type Invoker interface {
	// Call we have to implement a Call() method
	Call(interface{})
}

// Struct a structure
type Struct struct {
}

// Call Implementation of Call() method of Invoker
func (s *Struct) Call(p interface{}) {
	fmt.Println("from struct", p)
}

func main() {
	// state Invoker variable
	var invoker Invoker
	// initiate struct
	s := new(Struct)
	// assign struct to interface
	invoker = s
	// use interface call the method of struct
	invoker.Call("hello")

}
