package main

import (
	"fmt"
	"os"
)

// Go-Lang telah menyediakan banyak sekali package bawaan, salah satunya adalah package os
// Package os berisikan fungsionalitas untuk mengakses fitur sistem operasi secara independen (bisa digunakan  disemua sistem operasi)
// https://golang.org/pkg/os/

func main() {
	// os argumnets
	args := os.Args
	fmt.Println(args)

	// os hostname
	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname :", name)
	} else {
		fmt.Println("Error", err.Error())
	}

	// os environment variable
	username := os.Getenv("APP_USERNAME") // export APP_USERNAME=root
	password := os.Getenv("APP_PASSWORD") // export APP_PASSWORD=root

	if username == "" && password == "" {
		fmt.Println("Tidak ada environment variable")
	} else {
		fmt.Println("Username :", username)
		fmt.Println("Password :", password)
	}
}
