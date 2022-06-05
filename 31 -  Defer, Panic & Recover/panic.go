package main

import (
	"fmt"
)

func endApp() {
	fmt.Println("Aplikasi Selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APPLIKASI ERROR")
	}

	fmt.Println("APLIKASI BERJALAN")
}

func main() {
	runApp(true)
}
