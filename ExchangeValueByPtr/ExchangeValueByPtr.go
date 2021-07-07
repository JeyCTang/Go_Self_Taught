package main

//// swap function
//func swap(a, b *int) {
//	// get the value of pointer a, assign it to variable t
//	t := *a
//	// get the value of pointer b, assign it to the variable which a pointed to
//	*a = *b
//	// assign t to the variable which is appointed by b
//	*b = t
//}
//func main() {
//	// prepare two variables
//	x, y := 1, 2
//	// swap the value of two variables
//	swap(&x, &y)
//	// output the variables
//	fmt.Println(x, y)
//}

// If we swap two pointers
func swap(a, b *int) (*int, *int){
	b, a = a, b
	return a, b
}
func main() {
	x, y := 1, 2
	ptrX := &x
	ptrY := &y
	println(ptrX, ptrY)
	ptrX, ptrY = swap(ptrX, ptrY)
	println(ptrX, ptrY)
	//ptrY, ptrX = ptrX, ptrY
	//println(ptrX, ptrY)
	//println(*ptrX, *ptrY)
	//println(x, y)
	//println(x, &x, ptrX, y, &y, ptrY)
}
