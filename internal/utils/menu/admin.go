package menu

import (
	"clothing-pair-project/internal/utils/messages"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func AdminMenu(db *sqlx.DB, message string) {
	fmt.Println()
	fmt.Println("=====================================")
	fmt.Println("Admin Menu")
	fmt.Println("1. Manage User")
	fmt.Println("2. Manage Product")
	fmt.Println("3. Manage Collection")
	fmt.Println("4. Manage Category")
	fmt.Println("5. Manage Color")
	fmt.Println("6. Manage Size")
	fmt.Println("7. Manage Payment Method")
	fmt.Println("8. Reports")
	fmt.Println("0. Logout")
	fmt.Println("=====================================")

	messages.PrintMessage(message)

	var input string
	fmt.Print("Choose option: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		message = "No input entered"
		messages.PrintMessage(message)
		AdminMenu(db, message)
	}

	switch input {
	case "1":
		ManageUserMenu(db, "")
	case "2":
		ManageProductMenu(db, "")
	case "3":
		// TODO: manage collection menu
	case "4":
		// TODO: manage category menu
	case "5":
		// TODO: manage color menu
	case "6":
		// TODO: manage size menu
	case "7":
		// TODO: manage payment method menu
	case "8":
		// TODO: reports menu
	case "0":
		message = "Logging out..."
		messages.PrintMessage(message)
		DashboardMenu(db, message)
	default:
		message = "Invalid input"
		messages.PrintMessage(message)
		AdminMenu(db, message)
	}
}
