package main

import (
	"fmt"
	"reflect"
)

// Dalam bahasa pemrograman, biasanya ada fitur Reflection, dimana kita bisa melihat struktur kode kita pada saat aplikasi sedang berjalan
// Hal ini bisa dilakukan di Go-Lang dengan menggunakan package reflect
// Fitur ini mungkin tidak bisa dibahas secara lengkap dalam satu video, Anda bisa eksplorasi package reflec ini secara otodidak
// Reflection sangat berguna ketika kita ingin membuat library yang general sehingga mudah digunakan
// https://golang.org/pkg/reflect/

type Sample struct {
	Name string `required:"true" max:"10"` // Struct Tag
}

type Contoh struct {
	Name        string `required:"true"`
	Description string `required:"true"`
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" { // jika isi field kosong maka return false
				return false
			}
		}
	}
	return true
}

func main() {
	sample := Sample{"Hakim"}
	sampleType := reflect.TypeOf(sample)
	structField := sampleType.Field(0)
	required := structField.Tag.Get("required")

	fmt.Println(structField.Name)
	fmt.Println(required)

	sample.Name = ""
	fmt.Println(IsValid(sample)) // false krn isi field kosong

	contoh := Contoh{
		Name:        "Ecobag",
		Description: "Ok",
	}
	fmt.Println(IsValid(contoh))
}
