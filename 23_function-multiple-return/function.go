package main

import "fmt"

func getFullName() (string, string) {
	return "Ecobag", "Hakim"
}

func main() {
	firstName, _ := getFullName()
	fmt.Println(firstName)
}
