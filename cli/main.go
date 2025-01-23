package main

import (
	"clothing-pair-project/config"
	"clothing-pair-project/entity"
	"clothing-pair-project/handler"
	"clothing-pair-project/service"
	"log"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect db", err.Error())
	}
	defer db.Close()

	PaymentMethodHandler := handler.NewPaymentMethodHandler(db)
	paymentMethodService := service.NewPaymentMethodService(PaymentMethodHandler)

	addPaymentMethod := entity.PaymentMethod{
		Name: "Credit Card",
	}

	err = paymentMethodService.Add(addPaymentMethod)
	if err != nil {
		log.Fatal("Failed to add payment method:", err.Error())
	}

	log.Println("Successfully added payment method:", addPaymentMethod)

	paymentMethods, err := paymentMethodService.FindAll()
	if err != nil {
		log.Fatal("Failed to find all payment methods:", err.Error())
	}

	log.Println("Successfully found payment methods:", paymentMethods)
}
