package menu

import (
	"clothing-pair-project/internal/database/sqlrepo"
	"clothing-pair-project/internal/helper"
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

	passwordBytes, err := terminal.HidePassword("Enter Password:")
	if err != nil {
		fmt.Println("Error reading terminal:", err)
		return
	}
	password = string(passwordBytes)

	userRepository := sqlrepo.NewUserRepository(db)
	userService := services.NewUserService(userRepository)

	user, err := userService.GetUserByUsername(username)
	if err != nil || !helper.CheckPasswordHash(password, user.Password) {
		errorMessage := "Wrong username or password"
		fmt.Println(errorMessage)
		fmt.Println()

		DashboardMenu(db, errorMessage)
	} else if !user.Active {
		errorMessage := "User is inactive, please contact admin"
		fmt.Println(errorMessage)
		fmt.Println()

		DashboardMenu(db, errorMessage)
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
