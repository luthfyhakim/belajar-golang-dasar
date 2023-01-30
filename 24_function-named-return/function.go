package main

import "fmt"

func getFullName() (firstName, middleName, lastName string) {
	firstName = "Hakim"
	middleName = "Luthfy"
	lastName = "Ecobag"

	return
}

func main() {
	fmt.Println(getFullName())

	firstName, _, _ := getFullName()
	fmt.Println(firstName)
}
