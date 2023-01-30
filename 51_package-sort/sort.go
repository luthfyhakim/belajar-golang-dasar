package main

import (
	"fmt"
	"sort"
)

// Package sort adalah package yang berisikan utilitas untuk proses pengurutan
// Agar data kita bisa diurutkan, kita harus mengimplementasikan kontrak di interface sort.Interface
// https://golang.org/pkg/sort/

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	users := []User{
		{"Hakim", 18},
		{"Ecobag", 21},
		{"Evelyn", 17},
		{"Eve", 16},
	}

	fmt.Println("Before Sorting :", users) // sblm disorting

	sort.Sort(UserSlice(users)) // sorting

	fmt.Println("After Sorting :", users)
}
