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
		return errors.New("недостаточно средств")
	}
	user.Balance -= usd
	return nil
}

func main() {
	user := User{
		Name:    "Oleg",
		Balance: 10,
	}
	fmt.Println("Баланс до оплаты:", user.Balance)

	err := Pay(&user, 4)
	if err != nil {
		fmt.Println("Оплата не прошла:", err.Error())
	} else {
		fmt.Println("Оплата была произведена")
	}
	fmt.Println("Баланс после:", user.Balance)
}

// 2:07:00 end
