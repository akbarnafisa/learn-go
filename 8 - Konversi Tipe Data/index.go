package main

import "fmt"

func main() {
	var nilai32 int32 = 1231123123
	var nilai8 int8 = int8(nilai32)
	var nilai64 int64 = int64(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai8)
	fmt.Println(nilai64)

	var char = "e"
	fmt.Println(char[0])
	fmt.Println(string(char[0]))

}
