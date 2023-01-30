package main

import "fmt"

// Interface adalah tipe data Abstract, dia tidak memiliki implementasi langsung
// Sebuah interface berisikan definisi-definisi method
// Biasanya interface digunakan sebagai kontrak

type HasName interface {
	GetName() string
}

// kontrak dengan HasName
func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

// Setiap tipe data yang sesuai dengan kontrak interface, secara otomatis dianggap sebagai interface tersebut
// Sehingga kita tidak perlu mengimplementasikan interface secara manual
// Hal ini agak berbeda dengan bahasa pemrograman lain yang ketika membuat interface, kita harus menyebutkan secara eksplisit akan menggunakan interface mana

type Person struct { // contoh 1
	Name string
}

func (person Person) GetName() string { // implement method interface
	return person.Name
}

type Animal struct { // contoh 2
	Name string
}

func (animal Animal) GetName() string { // implement method interface
	return animal.Name
}

func main() {
	hakim := Person{
		Name: "Hakim",
	}
	sayHello(hakim)

	cat := Animal{
		Name: "Mozza",
	}
	sayHello(cat)
}
