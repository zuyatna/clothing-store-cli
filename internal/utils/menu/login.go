package menu

import (
	"clothing-pair-project/internal/utils/terminal"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func LoginMenu(db *sqlx.DB) {
	var username, password string

	fmt.Print("Enter username: ")
	_, err := fmt.Scanln(&username)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Print("Enter terminal: ")
	passwordBytes, err := terminal.HidePassword()
	if err != nil {
		fmt.Println("Error reading terminal:", err)
		return
	}
	password = string(passwordBytes)

	// TODO: Use the terminal

	fmt.Println(username, password)
	fmt.Println()
}
