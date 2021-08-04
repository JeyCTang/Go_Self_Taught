package main

import (
	"errors"
	"fmt"
)

// define the error when the divisor is 0
var errDivitionByZero = errors.New("division by zero")

func div(dividend, divisor int) (int, error) {
	// estimate if the divisor is 0 and return
	if divisor == 0 {
		return 0, errDivitionByZero
	}
	// execute the normal calculation process, return nil if it's wrong
	return dividend / divisor, nil
}
func main() {
	fmt.Println(div(1, 0))
}
