package main

import "fmt"

// Recursive function adalah function yang memanggil function dirinya sendiri
// Kadang dalam pekerjaan, kita sering menemui kasus dimana menggunakan recursive function lebih mudah dibandingkan tidak menggunakan recursive function
// Contoh kasus yang lebih mudah diselesaikan menggunakan recursive adalah Factorial

// contoh kasus factorial

func factorialLoop(number int) int {
	result := 1
	for i := number; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(number int) int {
	result := 1
	if number == 1 {
		return result
	} else {
		// terus memanggil itself sampai masuk if
		return number * factorialRecursive(number-1)
	}
}

func main() {
	result := factorialLoop(5)
	fmt.Println(result)
	fmt.Println(5 * 4 * 3 * 2 * 1)

	result = factorialRecursive(5)
	fmt.Println(result)
}
