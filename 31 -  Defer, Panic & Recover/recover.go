package main

import (
	"fmt"
)

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}
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
	runApp(false)
}
