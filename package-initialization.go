package main

import (
	_ "belajar-golang-dasar/blank-identifier"
	"belajar-golang-dasar/database"
	"fmt"
)

// Kadang kita hanya ingin menjalankan init function di package tanpa harus mengeksekusi salah satu function yang ada di package
// Secara default, Go-Lang akan komplen ketika ada package yang di import namun tidak digunakan
// Untuk menangani hal tersebut, kita bisa menggunakan blank identifier (_) sebelum nama package ketika melakukan import

func main() {
	result := database.GetDatabase()
	fmt.Println(result)

	fmt.Println(database.Application)
}
