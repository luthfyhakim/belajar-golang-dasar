package main

import "fmt"

func sayHelloTo(fisrtName string, lastName string) {
	fmt.Println("Hello to", fisrtName, lastName)
}

func main() {
	fisrtName := "Ecobag"
	sayHelloTo(fisrtName, "Hakim")
}
