package main

import "fmt"

type Person struct {
	Name, Address string
	Age           int
}

// struct method / struct function
func (person Person) sayHi(name string) {
	fmt.Println("Hello", name, "from", person.Name)
}

// function biasa
func sayHello(person Person, name string) {
	fmt.Println("Hello", name, "from", person.Name)
}

func main() {
	hakim := Person{
		Name: "Hakim",
	}

	hakim.sayHi("Ecobag")
	sayHello(hakim, "Ecobag")
}
