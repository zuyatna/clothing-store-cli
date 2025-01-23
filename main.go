package main

import (
	"clothing-pair-project/config"
	"fmt"
	"log"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect db", err.Error())
	}
	defer db.Close()

	for {
		dashboardMenu()
		var input int
		fmt.Print("Choose option: ")
		fmt.Scanln(&input)

		switch input {
		case 1:
			// TODO: login
		case 2:
			// TODO: register
		case 0:
			return
		default:
			fmt.Println("Invalid input")
		}
	}

	// // setup payment method service
	// PaymentMethodHandler := handler.NewPaymentMethodHandler(db)
	// paymentMethodService := service.NewPaymentMethodService(PaymentMethodHandler)

	// cli.AddPaymentMethod(*paymentMethodService, "Credit Card")
	// cli.FindAllPaymentMethods(*paymentMethodService)
	// cli.FindPaymentMethodByID(*paymentMethodService, 1)
	// cli.UpdatePaymentMethod(*paymentMethodService, 1, "Bank Transfer")
	// cli.DeletePaymentMethod(*paymentMethodService, 1)
}

func dashboardMenu() {
	fmt.Println("=====================================")
	fmt.Println("Dashboard Menu")
	fmt.Println("1. Login")
	fmt.Println("2. Register")
	fmt.Println("0. Exit")
	fmt.Println("=====================================")
}
