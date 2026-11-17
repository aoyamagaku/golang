package methods

import (
	"fmt"
	"math/rand/v2"
)

type Crypto struct{}

func NewCrypto() Crypto {
	return Crypto{}
}

func (c Crypto) Pay(usd int) int{
	fmt.Println("Оплата криптой")
	fmt.Println("Размер оплаты:", usd, "USDT")

	id := rand.Int()
	return id
}

func (c Crypto) Cancel(id int){
	fmt.Println("Отмена операции", id)
}