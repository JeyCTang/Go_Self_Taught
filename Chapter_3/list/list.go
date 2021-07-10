package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.PushFront(67)
	l.PushBack("canon")
	l.PushBack("hammer")
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
