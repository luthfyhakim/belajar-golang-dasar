package main

import "fmt"

// Type Assertions merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan
// Fitur ini sering sekali digunakan ketika kita bertemu dengan data interface kosong

func random() interface{} {
	return "Hakim"
}

func main() {
	result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is a string")
	case int:
		fmt.Println("Value", value, "is a int")
	case bool:
		fmt.Println("Value", value, "is a boolean")
	default:
		fmt.Println("Unknown type")
	}
}
