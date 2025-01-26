package menu

import (
	"clothing-pair-project/internal/database/sql"
	"clothing-pair-project/internal/services"
	"clothing-pair-project/internal/utils/terminal"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func LoginMenu(db *sqlx.DB, message string) {
	terminal.Clear()

	var username, password string
	fmt.Print("Enter username: ")
	_, err := fmt.Scanln(&username)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	passwordBytes, err := terminal.HidePassword("Enter Password:")
	if err != nil {
		fmt.Println("Error reading terminal:", err)
		return
	}
	password = string(passwordBytes)

	userRepository := sql.NewUserRepository(db)
	userService := services.NewUserService(userRepository)

	user, err := userService.GetUserByUsername(username)
	if err != nil || user.Password != password {
		message = "Wrong username or password"
		fmt.Println(message)
		fmt.Println()

		DashboardMenu(db, message)
	} else {
		fmt.Println()
		fmt.Println("=====================================")
		fmt.Printf("Hi, %s \n", user.Username)

		if user.Role == "admin" {
			AdminMenu(db, "")
		} else {
			// TODO: Implement user menu
		}
	}
}
