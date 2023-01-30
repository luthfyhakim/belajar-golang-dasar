package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Luthfy Hakim", "Hakim"))
	fmt.Println(strings.Contains("Luthfy Hakim", "Eco"))

	fmt.Println(strings.Split("Luthfy Hakim", " "))

	fmt.Println(strings.ToLower("Luthfy Hakim"))

	fmt.Println(strings.ToUpper("Luthfy Hakim"))

	fmt.Println(strings.Trim("    Luthfy Hakim     ", " "))

	fmt.Println(strings.ReplaceAll("Hakim Luthfy Hakim", "Hakim", "Ecobag"))
}
