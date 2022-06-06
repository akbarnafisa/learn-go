package helper

import "fmt"

// var version = 1 // cant be imported because the name start with lowercase letter
var Application = "Belajar Golang"

func SayHello(name string) {
	fmt.Println("Hello", name)
}

// func sayGoodbye(name string) { // cant be imported because the name start with lowercase letter
// 	fmt.Println("Goodbye", name)
// }
