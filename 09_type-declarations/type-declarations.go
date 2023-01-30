package main

import "fmt"

func main() {
	// type declations
	// membuat tipe data baru dr tipe data yg sudah ada
	type Name string
	type Married bool

	var username Name = "username"
	fmt.Println(username)

	var statusMarried Married = false
	fmt.Println(statusMarried)
}
