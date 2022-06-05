package main

import (
	"fmt"
	"strings"
)

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) string {
	filteredName := filter(name)
	return filteredName
}

func filter(data string) string {
	if strings.ToLower(data) == "anjay" {
		return "*****"
	}
	return data
}

func main() {
	filter1 := func(data string) string {
		if strings.ToLower(data) == "anjay" {
			return "*****"
		}
		return data
	}
	fmt.Println(sayHelloWithFilter("Boobas", filter))
	fmt.Println(sayHelloWithFilter("Anjay", filter1))
	fmt.Println(sayHelloWithFilter("Googas", func(data string) string {
		if strings.ToLower(data) == "anjay" {
			return "*****"
		}
		return data
	}))

}
