package main

import "fmt"

func anything(i interface{}) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "anything"
	}
}

func main() {
	kosong := anything("")
	fmt.Println(kosong)
}
