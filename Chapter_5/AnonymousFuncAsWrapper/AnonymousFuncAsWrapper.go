package main

import (
	"flag"
	"fmt"
)

var skillParam = flag.String("skill", "", "skill to perform")

func main() {
	flag.Parse()
	var skill = map[string]func(){
		"fire": func() {
			fmt.Println("chicken fire")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angel fly")
		},
	}
	if f, ok := skill[*skillParam]; ok {
		f()
		//fmt.Println(ok)
	} else {
		fmt.Println("skill not found")
		//fmt.Println(ok)
	}
}
