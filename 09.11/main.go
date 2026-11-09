package main

import (
	"fmt"
)

func main() {
	s := sub()

	fmt.Println(s)
}

func sub() string {
	subscribed := true

	if subscribed {
		return "You are subscribed"
	} else {
		return "You are NOT subscribed"
	}
}
