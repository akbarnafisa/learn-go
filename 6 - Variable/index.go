package main

import "fmt"

func main() {
	var name string
	name = "kucing"
	fmt.Println(name)

	name = "sussy"
	fmt.Println(name)

	var test int8 = 123
	test = 21
	fmt.Println(test)

	noneedvar := 123
	fmt.Println(noneedvar)

	var (
		firstName = "123"
		lastName  = "321"
	)

	fmt.Println(firstName, lastName)
}
