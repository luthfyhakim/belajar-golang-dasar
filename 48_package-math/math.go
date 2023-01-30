package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Round(1.3)) // Membulatkan float64 keatas atau kebawah, sesuai dengan yang paling dekat
	fmt.Println(math.Floor(1.7)) // Membulatkan float64 kebawah
	fmt.Println(math.Ceil(1.3))  // Membulatkan float64 keatas

	fmt.Println(math.Max(10, 20))
	fmt.Println(math.Min(10, 20))
}
