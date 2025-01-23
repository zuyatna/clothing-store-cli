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

	// add payment method
	addPaymentMethod := entity.PaymentMethod{
		Name: "Credit Card",
	}

	err = paymentMethodService.Add(addPaymentMethod)
	if err != nil {
		log.Fatal("Failed to add payment method:", err.Error())
	}
	log.Println("Successfully added payment method:", addPaymentMethod)

	// find all payment methods
	paymentMethods, err := paymentMethodService.FindAll()
	if err != nil {
		log.Fatal("Failed to find all payment methods:", err.Error())
	}

	log.Println("Successfully found payment methods:", paymentMethods)

	// find payment method by ID
	paymentMethod, err := paymentMethodService.FindByID(1)
	if err != nil {
		log.Fatal("Failed to find payment method by ID:", err.Error())
	}
	log.Println("Successfully found payment method by ID:", paymentMethod)

	// delete payment method
	updatePaymentMethod := entity.PaymentMethod{
		PaymentMethodID: 1,
		Name:            "Bank Transfer",
	}
	err = paymentMethodService.Update(updatePaymentMethod)
	if err != nil {
		log.Fatal("Failed to update payment method:", err.Error())
	}
	log.Println("Successfully updated payment method:", updatePaymentMethod)

	// delete payment method
	err = paymentMethodService.Delete(1)
	if err != nil {
		log.Fatal("Failed to delete payment method:", err.Error())
	}
	log.Println("Successfully deleted payment method with ID 1")

	// reset increment
	err = paymentMethodService.ResetIncrement()
	if err != nil {
		log.Fatal("Failed to reset increment:", err.Error())
	}
	log.Println("Successfully reset increment")
}
