package main

import "fmt"

func main() {
	// sama dengan penulisan variable
	const firstName string = "luthfy"
	const lastName = "hakim"

	fmt.Println(firstName)
	fmt.Println(lastName)

	const (
		name1 = "ecobag"
		name2 = name1
	)
	fmt.Println(name1)
	fmt.Println(name2)
}
