package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

// Package container/ring adalah implementasi struktur data circular list
// Circular list adalah struktur data ring, dimana diakhir element akan kembali ke element awal (HEAD)
// https://golang.org/pkg/container/ring/

func main() {
	data := ring.New(10)

	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.FormatInt(int64(i), 10)
		data = data.Next() // next ke element selanjutnya
	}

	// fmt.Println(*data) // tidak bisa yang kayak gini

	data.Do(func(value interface{}) { // untuk print data
		fmt.Println(value)
	})
}
