package main

import (
	"fmt"
	"strings"
)

// StringProcess input the strings and the chain of process function
func StringProcess(list []string, chain []func(string) string) {
	// traverse every string
	for index, str := range list {
		// the first string we are going to process in the chain
		result := str
		// traverse every function in the chain
		for _, proc := range chain {
			// input the str for processing, return the result as the input of the next function in the chain
			result = proc(result)
		}
		// save the all processed result
		list[index] = result
	}
}

// custom process function
func removePrefix(str string) string {
	return strings.TrimPrefix(str, "go")
}

func main() {
	// the string list we are going to process
	list := []string{
		"go scanner",
		"go parser",
		"go compiler",
		"go printer",
		"go formater",
	}
	// The function chain
	chain := []func(string) string{
		removePrefix,
		strings.TrimSpace,
		strings.ToUpper,
	}
	// process strings
	StringProcess(list, chain)
	// output the processed string
	for _, str := range list {
		fmt.Println(str)
	}
}
