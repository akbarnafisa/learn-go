package main

import (
	"fmt"
	"go-jaman-now/database"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
