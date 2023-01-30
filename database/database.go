package database

import "fmt"

// Di bahasa pemrograman lain, biasanya ada kata kunci yang bisa digunakan untuk menentukan access modifier terhadap suatu function atau variable
// Di Go-Lang, untuk menentukan access modifier, cukup dengan nama function atau variable
// Jika nama nya diawali dengan hurup besar, maka artinya bisa diakses dari package lain, jika dimulai dengan hurup kecil, artinya tidak bisa diakses dari package lain

var connection string

var version = 1                                 // tidak bisa diakses dr luar
var Application string = "Belajar Golang Dasar" // diawali huruf besar

// Saat kita membuat package, kita bisa membuat sebuah function yang akan diakses ketika package kita diakses
// Ini sangat cocok ketika contohnya, jika package kita berisi function-function untuk berkomunikasi dengan database, kita membuat function inisialisasi untuk membuka koneksi ke database
// Untuk membuat function yang diakses secara otomatis ketika package diakses, kita cukup membuat function dengan nama init

func init() { // package initialization
	fmt.Println("--------------------------------")
	fmt.Println("func init => otomatis dipanggil ketika package tersebut dipanggil")
	connection = "MySQL"
	fmt.Println(version)
	fmt.Println("--------------------------------")
}

func GetDatabase() string {
	return connection
}
