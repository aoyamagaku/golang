package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Input:")

	if ok := scanner.Scan(); !ok {
		fmt.Println("Input Error")
		return
	}

	text := scanner.Text()

	fields := strings.Fields(text)

	fmt.Println("Слова:", fields)
	fmt.Println("Команда:", fields[0])

	cmd := fields[0]
	if cmd == "добавить" {
		fmt.Println("вы хотите добавить", text)
	} else if cmd == "удалить" {
		fmt.Println("вы хотите удалить")
	}
}
