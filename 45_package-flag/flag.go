package main

import (
	"flag"
	"fmt"
)

// Package flag berisikan fungsionalitas untuk memparsing command line argument
// https://golang.org/pkg/flag/

func main() {
	// -host=localhost -user=admin -password=root -port=9090
	var host *string = flag.String("host", "localhost", "Put your name host") // name, default_value, helper/desc
	username := flag.String("user", "root", "Put your username")
	password := flag.String("password", "", "Put your password")
	port := flag.Int("port", 3306, "Put your port")

	flag.Parse() // harus dipanggil utk memparsing

	fmt.Println("Host :", *host) // hasil dlm bentuk pointer
	fmt.Println("Username :", *username)
	fmt.Println("Password :", *password)
	fmt.Println("Port :", *port)
}
