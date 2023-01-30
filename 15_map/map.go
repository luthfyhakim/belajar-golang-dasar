package main

import "fmt"

func main() {
	person := map[string]string{ // map[type_key]type_value
		"name":    "ecobag",
		"address": "jati",
	}

	fmt.Println(person)
	fmt.Println(person["name"])

	// membuat map baru
	book := make(map[string]string)
	book["title"] = "snowy"
	book["author"] = "ecobag"
	book["test"] = "test"
	fmt.Println(book)

	delete(book, "test")

	fmt.Println(book)
}
