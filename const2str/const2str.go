package main

import "fmt"

type ChipType int

const(
	None ChipType = iota
	CPU
	GPU
)

func (c ChipType) String() string {
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}
	return "N/A"
}

func main() {
	//output the value and format of CPU
	fmt.Printf("%s %d\n", CPU, GPU)
}
