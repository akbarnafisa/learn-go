package main

import "fmt"

func main() {
	type NoKTP string
	var name NoKTP = "12345"
	fmt.Println(name)
}
