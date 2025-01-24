package menu

import (
	"clothing-pair-project/handler"
	"clothing-pair-project/service"
	"fmt"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/olekukonko/tablewriter"
)

func ManagePaymentMethodMenu(db *sqlx.DB) {
	pmHandler := handler.NewPaymentMethodHandler(db)
	pmService := service.NewPaymentMethodService(pmHandler)
	choosePaymentMethod(pmService)
}

func allPaymentMethod(paymentService *service.PaymentMethodService) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name"})

	payments, err := paymentService.FindAll()
	if err != nil {
		fmt.Println("Error fetching users:", err)
		return
	}

	if len(payments) == 0 {
		fmt.Println("No product found.")
		return
	}

	for _, payment := range payments {
		table.Append([]string{
			strconv.Itoa(payment.PaymentMethodID),
			payment.Name,
		})
	}
	table.Render()

	fmt.Println()
}

func choosePaymentMethod(paymentService *service.PaymentMethodService) {
	fmt.Println()
	fmt.Println("=====================================")
	fmt.Println("Choose Payment")
	fmt.Println("=====================================")

	allPaymentMethod(paymentService)

	var input string
	fmt.Print("0. Back: ")
	fmt.Scanln(&input)
	if input == "0" {
		return
	}

	// paymentID, err := strconv.Atoi(input)
	// if err != nil {
	// 	fmt.Println("Invalid input, please enter a valid product ID.")
	// 	return
	// }

	// id := &paymentID
	// payments, err := paymentService.FindByID(id)
	// if err != nil {
	// 	fmt.Println("Invalid input")
	// 	return
	// }

	// if len(payment) == 0 {
	// 	fmt.Println("Invalid input")
	// }
}
