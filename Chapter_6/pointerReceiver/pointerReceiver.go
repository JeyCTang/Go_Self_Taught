package main

import "fmt"

// Property define struct
type Property struct {
	value int // value of property
}

// SetValue set value of property
func (p *Property) SetValue(v int) {
	// change the member value of property
	p.value = v
}

// Value get the value of property
func (p *Property) Value() int {
	return p.value
}

func main() {
	// initiate property
	p := new(Property)
	// set value
	p.SetValue(100)
	// print value
	fmt.Println(p.Value())
	fmt.Println(p.value)
}
