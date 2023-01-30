package main

import (
	"container/list"
	"fmt"
)

// Package container/list adalah implementasi struktur data double linked list di Go-Lang
// https://golang.org/pkg/container/list/

func main() {
	data := list.New()

	data.PushBack("Ecobag")
	data.PushBack("Luthfy")
	data.PushBack("Hakim")
	data.PushFront("Kubectl") // paling depan

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
