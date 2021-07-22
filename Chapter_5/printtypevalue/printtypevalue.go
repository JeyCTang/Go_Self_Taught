package main

import (
	"bytes"
	"fmt"
)

func printTypeValue(slist ...interface{}) string {
	// connecting chars
	var b bytes.Buffer
	// traverse variables
	for _, s := range slist {
		// convert interface to string
		str := fmt.Sprintf("%v", s)
		// description of string
		var typeString string
		// assert s
		switch s.(type) {
		case bool:
			typeString = "bool"
		case string:
			typeString = "string"
		case int:
			typeString = "int"
		}
		// head of value
		b.WriteString("value: ")
		// value
		b.WriteString(str)
		// head of type
		b.WriteString(" type: ")
		// type
		b.WriteString(typeString)
		// line break
		b.WriteString("\n")
	}
	return b.String()
}

func main() {
	// print various variables via printTypeValue function
	fmt.Println(printTypeValue(100, "str", true))
}
