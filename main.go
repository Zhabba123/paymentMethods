package main

import (
	"github.com/k0kubun/pp"
	"pay/payments"
	"pay/payments/methods"
)

func main() {
	method := methods.NewBank()

	paymentModule := payments.NewPaymentModule(method)

	paymentModule.Pay("Бургер", 5)
	idPhone := paymentModule.Pay("Телефон", 500)
	idGame := paymentModule.Pay("Игра", 20)

	paymentModule.Cancel(idPhone)

	allInfo := paymentModule.AllInfo()

	pp.Println("Все наши оплаты:", allInfo)

	gameInfo := paymentModule.Info(idGame)
	pp.Println("Game info:", gameInfo)
}
