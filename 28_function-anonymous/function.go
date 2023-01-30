package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blacklisted", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	// assign ke variable
	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("root", blacklist)

	// assign langsung ke parameter
	registerUser("admin", func(name string) bool {
		return name == "admin"
	})

	registerUser("root", func(name string) bool {
		return blacklist(name)
	})
}
