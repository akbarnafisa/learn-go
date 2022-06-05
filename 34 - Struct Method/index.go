package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHelloTo(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {
	joko := Customer{
		Name:    "Joko",
		Address: "Pantai Gading",
		Age:     59,
	}

	joko.sayHelloTo("Agus")

}
