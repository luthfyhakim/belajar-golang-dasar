package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	hakim := Man{
		Name: "hakim",
	}
	hakim.Married()

	fmt.Println(hakim.Name)
}
