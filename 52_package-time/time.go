package main

import (
	"fmt"
	"time"
)

// Package time adalah package yang berisikan fungsionalitas untuk management waktu di Go-Lang
// https://golang.org/pkg/time/

func main() {
	now := time.Now()
	fmt.Println(now.Local())

	utc := time.Date(2009, time.December, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(utc.Local())

	parse, _ := time.Parse(time.RFC3339, "2022-12-28")
	fmt.Println(parse)
}
