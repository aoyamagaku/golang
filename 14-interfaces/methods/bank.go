package methods

import (
	"fmt"
	"math/rand/v2"
)

type Bank struct{}

func NewBank() Bank {
	return Bank{}
}

func (c Bank) Pay(usd int) int {
	fmt.Println("Оплата картой")
	fmt.Println("Размер оплаты:", usd, "dollars")

	id := rand.Int()
	return id
}

func (c Bank) Cancel(id int) {
	fmt.Println("Отмена операции", id)
}
