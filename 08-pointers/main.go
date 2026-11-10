package main

import (
	"fmt"
)

func main() {
	number := 15

	var ptr *int = &number

	fmt.Println("number:", number)
	fmt.Println("ptr:", ptr)
	if ptr != nil {
		fmt.Println(*ptr)
	} else {
		fmt.Println("ptr is nil")
	}
}

// func changeName(s *string) {
// 	*s = "Ivan"
// }

// func foo(n *int) {
// 	// fmt.Println(n)
// 	fmt.Println(*n)

// 	*n = 10
// }

// func boo(n int) {
// 	n = 10
// }
