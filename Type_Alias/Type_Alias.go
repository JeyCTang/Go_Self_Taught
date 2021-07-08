package main

import (
	"fmt"
	"reflect"
)

// Brand define brand struct
type Brand struct {
}

// Show Add a Show() method for brand
func (t Brand) Show() {

}

// FakeBrand Define an alias name for Brand as FakeBrand
type FakeBrand = Brand

// Vehicle Define struct of vehicle
type Vehicle struct {
	// embedding Brand and FakeBrand struct
	FakeBrand
	Brand
}

func main() {
	// state variable a, the type is struct of Vehicle
	var a Vehicle
	// Call show() of FakeBrand
	a.FakeBrand.Show()
	// Get the type reflection object
	ta := reflect.TypeOf(a)
	// traverse all elements of a
	for i := 0; i < ta.NumField(); i++ {
		// element info of a
		f := ta.Field(i)
		// print the name and type of element
		fmt.Printf("Field Name: %v, Field Type: %v\n", f.Name, f.Type.Name())
	}
}
