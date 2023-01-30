package main

import (
	"fmt"
	"regexp"
)

// Package regexp adalah utilitas di Go-Lang untuk melakukan pencarian regular expression
// Regular expression di Go-Lang menggunakan library C yang dibuat Google bernama RE2
// https://github.com/google/re2/wiki/Syntax
// https://golang.org/pkg/regexp/

func main() {
	regex := regexp.MustCompile(`e([a-z])o`)

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("edo"))
	fmt.Println(regex.MatchString("eYo"))

	fmt.Println(regex.FindAllString("eko edo egi ego e1o eto", 10))
}
