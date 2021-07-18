package main

import "fmt"

// Accumulate provide a value, everytime we call the function it will accumulate the value
func Accumulate(value int) func() int {
	// return a closure
	return func() int {
		// accumulating
		value++
		// return accumulated value
		return value
	}
}

func main() {
	// create a accumulator, initiate value is 1
	accumulator := Accumulate(1)
	// accumulate 1, and print
	fmt.Println(accumulator())
	fmt.Println(accumulator())
	// print the address of accumulator
	fmt.Printf("%p\n", accumulator)
	// create the other accumulate, initiate value is 1
	accumulator2 := Accumulate(1)
	// accumulating and print out
	fmt.Println(accumulator2())
	// print out the address of accumulator2
	fmt.Printf("%p\n", accumulator2)

}
