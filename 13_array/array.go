package main

import "fmt"

func main() {
	// declarasi array dulu
	var names [3]string

	names[0] = "Aulia"
	names[1] = "Luthfi"
	names[2] = "Hakim"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// langsung masukkan value
	var values = [3]int{
		90,
		80,
		85,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
}
