package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// looping slice
	slices := []string{"Hakim", "Luthfy", "Ecobag"}

	for i := 0; i < len(slices); i++ {
		fmt.Println("Index ke", i, "=", slices[i])
	}

	// looping array or slice or map with for range
	// tanda _ digunakan saat tidak ingin Println variable agar tdk error
	for _, slice := range slices {
		fmt.Println(slice)
	}

	// map
	names := make(map[int]string)
	names[1] = "Pensil"
	names[2] = "Tas"
	names[3] = "Pulpen"
	names[4] = "Penggaris"

	for key, value := range names {
		fmt.Println(key, "=", value)
	}
}
