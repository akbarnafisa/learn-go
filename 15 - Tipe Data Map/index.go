package main

import "fmt"

func main() {
	person := map[string]string{
		"fistName": "Agus",
		"lastName": "Budi",
	}

	person["job"] = "kuli"
	person["age"] = "35"

	delete(person, "job")

	fmt.Println(person)
	fmt.Println(len(person))
	fmt.Println(person["Agus"])
	fmt.Println(person["fistName"])

	data := make(map[string]string)
	data["kucing"] = "meong"

	fmt.Println(data)

}
