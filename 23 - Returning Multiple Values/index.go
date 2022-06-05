package main

import "fmt"

func kocheng(firstName, lastName string) (string, string) {
	firstName = "Default"
	lastName = "Default"
	return firstName, lastName
}

func main() {
	firstName, _ := kocheng("Kocheng", "Gawrong")
	fmt.Println(firstName)
}
