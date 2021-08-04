package main

import (
	"fmt"
	"runtime"
)

// the contents we have to transfer when panic happens
type panicContext struct {
	function string // the function
}

//
func ProtectRun(entry func()) {
	// defer function
	defer func() {
		// when panic happen, get the contents transferred from panic and print it
		err := recover()

		switch err.(type) {
		case runtime.Error: // running error
			fmt.Println("runtime error:", err)
		default:
			fmt.Println("error", err)
		}
	}()
	entry()
}
func main() {
	fmt.Println("before running")
	// allow the manually trigger error
	ProtectRun(func() {
		fmt.Println("before manually panic")
		// transfer contexts via panic
		panic(&panicContext{
			"manually trigger panic",
		})
		fmt.Println("after manually panic")
	})
	// null pointer call error
	ProtectRun(func() {
		fmt.Println("before assign error")
		var a *int
		*a = 1
		fmt.Println("after assign error")
	})
	fmt.Println("after running")
}
