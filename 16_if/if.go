package main

import "fmt"

func main() {
	name := "Hakim"

	if name == "Hakim" {
		fmt.Println("Halo", name)
	} else if name == "Ecobag" {
		fmt.Println("Halo", name)
	} else {
		fmt.Println("Hallo Everyone!")
	}

	// if short statement
	if length := len(name); length > 5 {
		fmt.Println("Nama Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}
}
