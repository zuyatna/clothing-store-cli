package menu

import (
	"bufio"
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
	"strings"
)

func DashboardMenu(db *sqlx.DB) {
	for {
		fmt.Println("Welcome to the clothing store!")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("0. Exit")

		var input string
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Choose an option: ")
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			// TODO: Implement login
		case "2":
			// TODO: Implement register
		case "0":
			return
		default:
			fmt.Println("Invalid input. Please try again.")
		}
		fmt.Println()
	}
}
