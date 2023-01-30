package main

import (
	"belajar-golang-dasar/43_package-import/helper"
	"belajar-golang-dasar/43_package-import/other"
	"fmt"
)

func main() {
	result := helper.SayHelloFromHelper("Hakim")
	fmt.Println(result)

	other.SayHelloFromOther("Hakim")
}
