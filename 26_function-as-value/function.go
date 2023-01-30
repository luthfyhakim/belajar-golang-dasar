package main

import "fmt"

// function adalah first class citizen pada golang
// function juga merupakan tipe data, dan bisa disimpan di dalam variable

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	sayGoodBye := getGoodBye // function sbg value dr variable

	fmt.Println(sayGoodBye("Hakim"))
}
