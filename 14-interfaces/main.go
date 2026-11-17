package main

import (
	"fmt"

	"github.com/aoyamagaku/golang/14-interfaces/methods"
	"github.com/aoyamagaku/golang/14-interfaces/payments"
)

func main() {
	method := methods.NewBank()

	paymentModule := payments.NewPaymentModule(method)

	paymentModule.Pay("Burger", 5)
	idPhone := paymentModule.Pay("Phone", 500)
	paymentModule.Pay("Steam", 20)

	paymentModule.Cancel(idPhone)

	allInfo := paymentModule.AllInfo()

	fmt.Println("Payments:", allInfo)
}
