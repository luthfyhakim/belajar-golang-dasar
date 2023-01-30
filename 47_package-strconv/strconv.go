package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true") // string => boolean
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	int, _ := strconv.ParseInt("1000000", 10, 64) // string => int
	fmt.Println(int)

	intBinary, _ := strconv.ParseInt("1000000", 2, 64) // value, base(2 => binary, 8, 10 => decimal, 16), bit_size
	fmt.Println(intBinary)

	string := strconv.FormatInt(1000000, 2)
	fmt.Println(string) // hasilnya binary

	valueInt, _ := strconv.Atoi("2000000") // sama dengan ParseInt
	fmt.Println(valueInt)
}
