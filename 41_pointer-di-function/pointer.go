package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address := Address{"Karangan", "Jatim", ""}

	changeCountryToIndonesia(&address)

	fmt.Println(address)
}
