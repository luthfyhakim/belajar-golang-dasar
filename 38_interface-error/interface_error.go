package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai, pembagi int) (int, error) { // interface error
	if pembagi == 0 {
		return 0, errors.New("pembagi tidak boleh 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {
	// contohError := errors.New("contoh error")
	// fmt.Println(contohError)

	result, err := pembagian(100, 0) // syntak utk menangani error pd golang
	if err == nil {
		fmt.Println("Hasil =", result)
	} else {
		fmt.Println("Error =", err.Error())
	}
}
