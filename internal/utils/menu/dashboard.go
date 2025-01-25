package menu

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

func DashboardMenu(db *sqlx.DB) {
	for {
		fmt.Println("Welcome to the clothing store!")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("0. Exit")

		var input string
		fmt.Print("Choose an option: ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		fmt.Println()

		switch input {
		case "1":
			LoginMenu(db)
		case "2":
			// TODO: Implement register
		case "0":
			return
		default:
			fmt.Println("Invalid input. Please try again.")
		}
	}
}
