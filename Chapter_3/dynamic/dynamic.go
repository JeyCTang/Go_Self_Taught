package main

import "fmt"

func main() {
	// make([]T, size, cap)
	// cap is the number of pre-allocate element to the slice
	a := make([]int, 2)
	b := make([]int, 2, 10)

	fmt.Println(a, b)
	fmt.Println(len(a), len(b))

	// append() extend the memory size of slice
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("len:    %d    cap:    %d    pointer:%p\n", len(numbers), cap(numbers), numbers)
	}
	fmt.Println(numbers)

	var car []string
	// add the one element
	car = append(car, "Old Driver")
	// add more than one elements
	car = append(car, "Ice", "Sniper", "Monk")
	// add slice
	team := []string{"Pig", "Flyingcake", "Chicken"}
	car = append(car, team...)

	fmt.Println(car)

}
