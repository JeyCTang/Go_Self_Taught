package main

import (
	"fmt"
	"reflect"
)

// MyInt define int type to MyInt type
type MyInt int

// IsZero add IsZero() method for MyInt
func (m MyInt) IsZero() bool {
	return m == 0
}

// Add add Add() method for MyInt
func (m MyInt) Add(other int) int {
	return other + int(m)
}

func main() {
	var b MyInt
	b = 100
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(b.IsZero())
	b = 1
	fmt.Println(b.Add(2))
}
