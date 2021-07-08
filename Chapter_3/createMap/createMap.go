package main

import "fmt"

func main() {
	scene := make(map[string]int)
	scene["route"] = 66
	fmt.Println(scene["route"])

	v := scene["route2"]
	fmt.Println(v)

	v1, ok := scene["route"]
	v2, result := scene["haha"]
	fmt.Println(v1, ok)
	fmt.Println(v2, result)

	m := map[string]string{
		"W": "forward",
		"A": "left",
		"D": "right",
		"S": "backward",
	}
	fmt.Println(m)
}
