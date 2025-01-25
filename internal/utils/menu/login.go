package menu

import (
	"fmt"
	"syscall"

	"github.com/jmoiron/sqlx"
	"golang.org/x/term"
)

func LoginMenu(db *sqlx.DB) {
	var username, password string

	fmt.Print("Enter username: ")
	_, err := fmt.Scanln(&username)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Print("Enter password: ")
	passwordBytes, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println("\nFailed to read password:")
		fmt.Println(err.Error())
		return
	}
	password = string(passwordBytes)

	// TODO: Use the password

	fmt.Println(username, password)
	fmt.Println()
}
