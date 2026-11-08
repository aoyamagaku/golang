package main

import "fmt"

// 1 -- how functions are working in golang

func main() {
	fmt.Println("Before calling hello function")
	hello()
	fmt.Println("Aftre calling hello function")
}

func hello() {
	fmt.Println("Hello1")
	fmt.Println("Hello2")
}

//End at 6:08:00
