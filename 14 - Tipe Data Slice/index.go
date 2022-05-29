package main

import "fmt"

func main() {
	// var values = [...]int{
	// 	1,
	// 	2,
	// 	3,
	// 	4,
	// 	5,
	// 	6,
	// }

	// var gg = values[0:3]
	// var gg1 = values[2:]
	// var gg2 = values[:2]
	// var gg3 = values[:]

	// fmt.Println(gg)
	// fmt.Println(len(gg))
	// fmt.Println(cap(gg))

	// Becarefull!, slice is pass by reference
	// gg[0] = 999
	// values[1] = 666
	// fmt.Println(gg)

	// slice1 := values[:]
	// slice2 := append(slice1, 69)
	// slice2[0] = 333

	// fmt.Println(values)
	// fmt.Println(slice1)
	// fmt.Println(slice2)

	// fmt.Println(gg1)
	// fmt.Println(gg2)
	// fmt.Println(gg3)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "kucing"
	newSlice[1] = "garong"

	// fmt.Println(newSlice)
	// fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)
	fmt.Println(len(copySlice))

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
