package main

import "fmt"

func main() {
	seq := []string{"a", "b", "c", "d", "e"}
	// indicate the element index which will be deleted
	index := 2
	// check the elements before and after the deleted element
	fmt.Println(seq[:index], seq[index+1:])
	// connect the elements before and after the deleted element
	seq = append(seq[:index], seq[index+1:]...)

	fmt.Println(seq)
}
