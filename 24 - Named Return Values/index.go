package main

import "fmt"

func kocheng() (firstName string, lastName string) {
	firstName = "Default"
	lastName = "Default"

	return
}

func main() {
	// a, b, c := getFullName2()
	a, b := kocheng()

	fmt.Println(a)
	fmt.Println(b)
}
