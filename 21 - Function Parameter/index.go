package main

import "fmt"

func kocheng(firstName, lastName string) string {
	fmt.Println(firstName, lastName)
	return firstName
}

func gg(names []int) []int {
	return names
}

func main() {
	var values = [](int){
		1,
		2,
		3,
	}
	gg(values)
	kocheng("Kocheng", "Gawrong")
}
