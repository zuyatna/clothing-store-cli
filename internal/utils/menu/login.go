package menu

import (
	"clothing-pair-project/internal/database/sql"
	"clothing-pair-project/internal/services"
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

	passwordBytes, err := terminal.HidePassword()
	if err != nil {
		fmt.Println("Error reading terminal:", err)
		return
	}
	password = string(passwordBytes)

	userRepository := sql.NewUserRepository(db)
	userService := services.NewUserService(userRepository)

	user, err := userService.GetUserByUsername(username)
	if err != nil || user.Password != password {
		fmt.Println("Wrong username or password")
		fmt.Println()

		DashboardMenu(db)
	} else {
		fmt.Println()
		fmt.Println("=====================================")
		fmt.Printf("Hi, %s \n", user.Username)
		fmt.Println("=====================================")

		if user.Role == "admin" {
			// TODO: Implement admin menu
		} else {
			// TODO: Implement user menu
		}
	}
}
