package main

import "fmt"

func main() {
	// Tipe data Slice adalah potongan dari data Array
	months := [...]string{
		"Jan", // 0
		"Feb", // 1
		"Mar", // 2
		"Apr", // 3
		"May", // 4
		"Jun", // 5
		"Jul", // 6
		"Aug", // 7
		"Sep", // 8
		"Oct", // 9
		"Nov", // 10
		"Dec", // 11
	}

	/*
			1. Tipe Data Slice memiliki 3 data, yaitu pointer, length dan capacity
		 	2. Pointer adalah penunjuk data pertama di array para slice
			3. Length adalah panjang dari slice, dan
			4. Capacity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capacity
	*/
	/*
		array[low:high]
		array[low:]
		array[:high]
		array[:]
	*/
	slice1 := months[4:7] // Membuat slice dari array dimulai index low sampai index sebelum high
	slice2 := months[4:]  // Membuat slice dari array dimulai index low sampai index akhir di array
	slice3 := months[:7]  // Membuat slice dari array dimulai index 0 sampai index sebelum high
	slice4 := months[:]   // Membuat slice dari array dimulai index 0 sampai index akhir di array

	fmt.Println(slice1)
	fmt.Println(len(slice1)) // utk mendapatkan panjang
	fmt.Println(cap(slice1)) // utk mendapatkan capasitas

	fmt.Println(slice2)

	fmt.Println(slice3)

	fmt.Println(slice4)

	days := [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Miggu",
	}
	// Membuat slice baru dengan menambah data ke posisi terakhir slice
	// jika kapasitas sudah penuh, maka akan membuat array baru
	daySlice1 := days[4:]
	daySlice1[1] = "luthfy"
	fmt.Println(daySlice1)
	fmt.Println(days)

	// make slice
	newSlice := make([]string, 2, 5) // make([]type, length, capasity)
	newSlice[0] = "luthfy"
	newSlice[1] = "hakim"

	fmt.Println(newSlice)

	// copy slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(toSlice)

	// catatan
	iniArray := [...]int{1, 2, 3}
	iniSlice := []int{1, 2, 3}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
