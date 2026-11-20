package main

import (
	"errors"
	"fmt"
)

type User struct {
	Name    string
	Balance int
}

func Pay(user *User, usd int) error {
	if user.Balance-usd < 0 {
		return errors.New("Недостаточно средств")
	}
	return nil
}

func main() {
	user := User{
		Name:    "Oleg",
		Balance: 10,
	}

	err := Pay(&user, 54)

	if err != nil {
		fmt.Println("Оплата не прошла", err.Error())
	} else {
		fmt.Println("Оплата была произведена")
	}
}

// 2:07:00 end
