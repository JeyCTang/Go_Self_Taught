package main

import (
	"fmt"
	"math"
)

func main() {
	// output the range of various type of values
	fmt.Println("int8 range:", math.MinInt8, math.MaxInt8)
	fmt.Println("int16 range:", math.MinInt16, math.MaxInt16)
	fmt.Println("int32 range:", math.MinInt32, math.MaxInt32)
	fmt.Println("int64 range:", math.MinInt64, math.MaxInt64)

	// initiate a 32-bit integer
	var a int32 = 1047483647
	// output this variable with hex and decimal
	fmt.Printf("int32: 0x%x %d\n\n\n", a, a)

	// convert a to hex, it will cause block of numbers
	b := int16(a)
	// output this variable with hex and decimal
	fmt.Printf("int16: 0x%x %d\n\n\n", b, b)

	// save the constant with float-32 type
	var c float32 = math.Pi
	// convert it to int type, the number after decimal will be lose
	fmt.Println(int(c))
}
