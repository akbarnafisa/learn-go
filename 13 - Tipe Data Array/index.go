package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "1"
	names[1] = "2"
	names[2] = "3"

	// var gg [3]number = [1 2 3]

	var values = [3]int{
		1,
		2,
		3,
	}

	fmt.Println(names)
	fmt.Println(values)

}
