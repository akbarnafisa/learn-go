package main

import (
	"fmt"
)

func login() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	defer login()
	fmt.Println("Run Application")
}

func main() {
	runApplication()
}
