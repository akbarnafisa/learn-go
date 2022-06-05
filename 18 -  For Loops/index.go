package main

import "fmt"

func main() {
	// count := 10
	// for i := 0; i < count; i++ {
	// 	fmt.Println(i)
	// }

	// count := 0
	// for count <= 10 {
	// 	fmt.Println(count)
	// 	count++
	// }

	// count := 10
	// for i := 0; i < count; i++ {
	// 	fmt.Println(i)
	// }

	// slice := []string{"a", "b", "c", "d"}
	// for i, value := range slice {
	// 	fmt.Println("Index", i, " = ", value)
	// }

	slice := []string{"a", "b", "c", "d"}
	for _, value := range slice {
		fmt.Println(value)
	}
}
