package main

import "fmt"

func main() {
	var team0 [3]string
	team0[0] = "hammer"
	team0[1] = "soldier"
	team0[2] = "mum"

	fmt.Println(team0)

	var team1 = [3]string{"hammer", "soldier", "mum"}
	fmt.Println(team1)

	var team2 = [...]string{"hammer", "soldier", "mum"}
	fmt.Println(team2)

	for k, v := range team1 {
		fmt.Println(k, v)
	}
}
