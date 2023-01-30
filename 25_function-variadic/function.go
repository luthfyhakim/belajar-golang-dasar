package main

import "fmt"

// variable argumen

/*
1. Parameter yang berada di posisi terakhir, memiliki kemampuan dijadikan sebuah varargs
2. Varargs artinya datanya bisa menerima lebih dari satu input, atau anggap saja semacam Array.
3. Apa bedanya dengan parameter biasa dengan tipe data Array?
  - Jika parameter tipe Array, kita wajib membuat array terlebih dahulu sebelum mengirimkan ke function
  - JIka parameter menggunakan varargs, kita bisa langsung mengirim data nya, jika lebih dari satu, cukup gunakan tanda koma
*/

func sumAll(numbers ...int) int {
	total := 0
	for _, v := range numbers {
		total += v
	}
	return total
}

func main() {
	total := sumAll(10, 30, 60, 20, 70, 40)
	fmt.Println(total)

	slice := []int{10, 20, 60, 40, 50}
	total = sumAll(slice...)
	fmt.Println(total)
}
