package main

import "fmt"

// the actual function which will print contents
func rawPrint(rawList ...interface{}) {
	// traverse the variables list
	for _, a := range rawList {
		// print variables
		fmt.Println(a)
	}
}

func print1(slist ...interface{}) {
	// transfer slist to rawPrint
	rawPrint(slist...)
}
func main() {
	print1(1, 2, 3)
}
