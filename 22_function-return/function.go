package main

import "fmt"

func getName(name string) string {
	if name == "" {
		return "Hello Folks!"
	} else {
		return "Hello " + name
	}
	// return "Hello" + name
}

func main() {
	// fungsi yg mengembalikan data / return dpt disimpan pd variable
	result := getName("Ecobag")
	fmt.Println(result)

	fmt.Println(getName(""))
}
