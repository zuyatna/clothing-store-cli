package menu

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func DashboardMenu(db *sqlx.DB, message string) {
	fmt.Println("Welcome to the clothing store!")
	fmt.Println("1. Login")
	fmt.Println("2. Register")
	fmt.Println("0. Exit")

	fmt.Println()
	fmt.Println(message)
	fmt.Println()

	var input string
	fmt.Print("Choose option: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	switch input {
	case "1":
		LoginMenu(db)
	case "2":
		// TODO: Implement register
	case "0":
		fmt.Println("Goodbye!")
		err := db.Close()
		if err != nil {
			return
		}
		return
	default:
		message = "Invalid input. Please try again."
		DashboardMenu(db, message)
	}
}
