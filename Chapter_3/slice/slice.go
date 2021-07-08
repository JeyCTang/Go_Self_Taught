package main

import "fmt"

func main() {
	var a = [3]int{1, 2, 3}
	fmt.Println(a, a[1:2], a[1:])

	var a_1 = []int{1, 2, 3}
	fmt.Printf("a type: %T\n", a)
	fmt.Printf("a_1 type: %T\n", a_1)
	fmt.Println(a_1, a_1[:])

	// state string slice
	var strList []string
	// state number slice
	var numList []int
	// state empty slice
	var numListEmpty = []int{}
	// output these 3 slices
	fmt.Println(strList, numList, numListEmpty)
	// output the length of 3 slices
	fmt.Println(len(strList), len(numList), len(numListEmpty))
	// predicate if slice is null
	fmt.Println(strList == nil)
	fmt.Println(numList == nil)
	fmt.Println(numListEmpty == nil)
}
