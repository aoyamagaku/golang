package methods

import (
	"fmt"
	"math/rand/v2"
)

type PayPal struct{}

func NewPayPal() PayPal {
	return PayPal{}
}

func (c PayPal) Pay(usd int) int {
	fmt.Println("Оплата PayPal")
	fmt.Println("Размер оплаты:", usd, "USD")

	id := rand.Int()
	return id
}

func (c PayPal) Cancel(id int) {
	fmt.Println("Отмена операции", id)
}
