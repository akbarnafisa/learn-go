package main

import "fmt"

func main() {
	name := "s"
	switch name {
	case "Booba":
		fmt.Println("1")
	case "Agus":
		fmt.Println("Amogus")
	default:
		fmt.Println("GGWP")
	}

	number := 1
	switch {
	case number > 1:
		fmt.Println("1")
	case number > 5:
		fmt.Println("Amogus")
	default:
		fmt.Println("GGWP")
	}
}
