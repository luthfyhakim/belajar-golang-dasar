package main

import "fmt"

// struct itu seperti class pd OOP

// Struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
// Struct biasanya representasi data dalam program aplikasi yang kita buat
// Data di struct disimpan dalam field
// Sederhananya struct adalah kumpulan dari field

// Struct adalah template data atau prototype data
// Struct tidak bisa langsung digunakan
// Namun kita bisa membuat data/object dari struct yang telah kita buat

type Customer struct { // name struct UperCase
	Name, Address string
	Age           int
}

func main() {
	// cara 1
	var ecobag Customer
	ecobag.Name = "Ecobag"
	ecobag.Address = "Jakarta"
	ecobag.Age = 16

	fmt.Println(ecobag)
	fmt.Println(ecobag.Name)
	fmt.Println(ecobag.Address)
	fmt.Println(ecobag.Age)

	// cara 2
	// struct literals
	hakim := Customer{
		Name:    "Hakim",
		Address: "Surabaya",
		Age:     18,
	}
	fmt.Println(hakim)

	// cara 3 (tidak disarankan)
	eco := Customer{"Eco", "Jati", 13}
	fmt.Println(eco)
}
