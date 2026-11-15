package main

import "fmt"

//	func main() {
//		intSlice := make([]int, 0)
//
//		intSlice = append(intSlice, 10, 20, 30)
//
//		fmt.Println(intSlice)
//	}
func main() {
	crime := map[string]bool{
		"Max":  true,
		"Jack": true,
		"Mike": false,
	}
	value, ok := crime["Mike"]

	if !ok {
		fmt.Println("Человека нет в базе")
		return
	}
	fmt.Println("Человек найден в базе")

	if value {
		fmt.Println("Человек судим")
		return
	}
	fmt.Println("Человек не судим")
}
