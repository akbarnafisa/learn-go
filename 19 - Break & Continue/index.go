package main

import "fmt"

func main() {
	// count := 10
	// for i := 0; i < count; i++ {
	// 	if i == 5 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	count := 10
	for i := 0; i < count; i++ {
		if i%2 == 0 {
			break
		}
		fmt.Println(i)
	}
}
