package main

import "fmt"

// Data define a structure which will be used to test the transportation of deliver
type Data struct {
	complex  []int      // to test slice
	instance InnerData  // InnerData is also a struct
	ptr      *InnerData // the pointer of InnerData
}
type InnerData struct {
	a int
}

// test function
func passByValue(inFunc Data) Data {
	// output the members of values
	fmt.Printf("inFunc value: %+v\n", inFunc)
	// output the pointer of inFunc
	fmt.Printf("inFunc ptr: %p\n", &inFunc)

	return inFunc
}

// test process
func main() {
	in := Data{
		complex:  []int{1, 2, 3},
		instance: InnerData{5},
		ptr:      &InnerData{1},
	}
	// input the members of structure
	fmt.Printf("in value: %+v\n", in)
	// input the pointer of structure
	fmt.Printf("in ptr: %p\n", &in)

	// input the value to structure, we will get same type struct return
	out := passByValue(in)
	// output the members of structure
	fmt.Printf("out value: %+v\n", out)
	// output the pointer of structure
	fmt.Printf("out ptr: %p\n", &out)
}
