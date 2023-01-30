package main

import "fmt"

func main() {
	var name string // penulisan variable versi satu
	name = "hakim"
	fmt.Println(name)

	name = "luthfy"
	fmt.Println(name)

	var age = 25 // penulisan variable versi dua
	fmt.Println(age)

	// penulisan variable versi tiga (tanpa tanda var)
	country := "indonesia" // tanda := hanya wajib ketika inisialisasi pertama
	fmt.Println(country)

	country = "america"
	fmt.Println(country)

	// penulisan variable versi empat
	var (
		firstName = "luthfy"
		lastname  = "hakim"
	)
	fmt.Println(firstName)
	fmt.Println(lastname)
}
