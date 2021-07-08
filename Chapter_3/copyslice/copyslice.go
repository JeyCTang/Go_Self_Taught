package main

import "fmt"

func main() {
	// set the number to 1000
	const elementCount = 1000
	// pre-allocate enough element to slice
	srcData := make([]int, elementCount)
	// assign value to slice
	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}
	// refer slice data
	refData := srcData
	// pre-allocate enough element to slice
	copyData := make([]int, elementCount)
	// copy data to the new slice
	copy(copyData, srcData)
	// change the first element of original slice
	srcData[0] = 999
	// print the first element of refData
	fmt.Println(refData[0])
	// print the first and last element of copyData
	fmt.Println(copyData[0], copyData[elementCount-1])
	// copy the original data from index 4 - 6(not include)
	copy(copyData, srcData[4:6])
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", copyData[i])
	}
}
