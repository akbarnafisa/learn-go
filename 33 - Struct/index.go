package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var agus Customer
	agus.Name = "agus"
	agus.Address = "Jakarta"
	agus.Age = 30

	joko := Customer{
		Name:    "Joko",
		Address: "Pantai Gading",
		Age:     59,
	}

	budi := Customer{
		"Budi",
		"Indonesia",
		59,
	}

	fmt.Println(agus)
	fmt.Println(joko)
	fmt.Println(budi)

}
