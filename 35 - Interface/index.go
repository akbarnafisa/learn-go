package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name    string
	Address string
}

func (person Person) GetName() string {
	return person.Name
}

func main() {
	var kucing Person
	kucing.Name = "meong"
	SayHello(kucing)
}
