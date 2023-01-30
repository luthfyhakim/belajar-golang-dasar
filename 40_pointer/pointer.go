package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{
		City:     "London",
		Province: "Java",
		Country:  "United Kingdom",
	}

	// var address4 *Address = &Address{"London", "Java", "United Kingdom"}

	address2 := &address1 // pointer
	address3 := &address1

	address2.City = "Surabaya"

	*address2 = Address{"Jati", "Karangan", "Trenggalek"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *Address = new(Address) // membuat pointer baru
	address4.City = "Jakarta"
	fmt.Println(address4)
}
