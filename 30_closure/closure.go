package main

import "fmt"

// Closure adalah kemampuan sebuah function berinteraksi dengan data-data disekitarnya dalam scope yang sama
// Harap gunakan fitur closure ini dengan bijak saat kita membuat aplikasi

// hanya bisa mengakses data yang di atasnya (scope)

func main() {
	name := "Ecobag"
	counter := 0

	increment := func() {
		name := "Hakim"
		fmt.Println("Increment")
		fmt.Println(name)
		counter++ // closure
	}

	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
