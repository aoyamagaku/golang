package main

import (
	"fmt"
)

func main() {
	fmt.Println("Main start")
	defer func() {
		fmt.Println("Main end")
	}()

	// fmt.Println("Hello 1")
	// fmt.Println("Hello 2")
	// hello()

	foo()

}

func hello() {
	fmt.Println("Hello start")
	defer func() {
		fmt.Println("Hello end")
	}()

	fmt.Println("Hello 1")
	fmt.Println("Hello 2")
}

func foo() {
	fmt.Println("foo")
	defer func() {
		fmt.Println("First defer")
	}()

	defer func() {
		fmt.Println("Second defer")
	}()

	defer func() {
		fmt.Println("Third defer")
	}()
}
