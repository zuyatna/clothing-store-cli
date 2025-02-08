package menu

import (
	"clothing-pair-project/internal/database/sqlrepo"
	"clothing-pair-project/internal/helper"
	"clothing-pair-project/internal/services"
	"clothing-pair-project/internal/utils/key_input"
	"clothing-pair-project/internal/utils/terminal"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func LoginMenu(db *sqlx.DB) {
	var username, password string

	username, err := key_input.Username()
	if err != nil {
		fmt.Println("Error reading username:", err)
	}

	passwordBytes, err := terminal.HidePassword("Enter Password:")
	if err != nil {
		fmt.Println("Error reading terminal:", err)
		return
	}
	password = string(passwordBytes)

	userRepository := sqlrepo.NewUserQuery(db)
	userService := services.NewUserService(userRepository)

	user, err := userService.GetUserByUsername(username)
	if err != nil || !helper.CheckPasswordHash(password, user.Password) {
		errMsg := "Wrong username or password"
		fmt.Println(errMsg)
		DashboardMenu(db, errMsg)
	} else if !user.Active {
		errMsg := "User is inactive, please contact admin"
		fmt.Println(errMsg)
		DashboardMenu(db, errMsg)
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
