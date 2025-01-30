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
	fmt.Println("1. Manage Users")
	fmt.Println("2. Manage Products")
	fmt.Println("3. Manage Categories")
	fmt.Println("4. Manage Colors")
	fmt.Println("5. Manage Sizes")
	fmt.Println("6. Manage Payment Methods")
	fmt.Println("7. Reports")
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
		ManageCategoryMenu(db, "")
	case "4":
		// TODO: manage color menu
	case "5":
		// TODO: manage size menu
	case "6":
		// TODO: manage payment method menu
	case "7":
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
