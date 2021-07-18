package main

import "fmt"

func main() {
	// traverse array， slice， get the index and value
	for key, value := range []int{1, 2, 3, 4} {
		fmt.Printf("key: %d  value: %d\n", key, value)
	}
	// traverse strings, get the char
	var str = "hello world!"
	for key, value := range str {
		fmt.Printf("key: %d, value:%c\n", key, value)
	}
	// traverse map, you will get key and value
	m := map[string]int{
		"hello": 100,
		"world": 200,
	}
	for key, value := range m {
		fmt.Println(key, value)
	}
	// traverse channel, you will only receive one value in the channel
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
}
