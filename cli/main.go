package main

import (
	"clothing-pair-project/cli/menu"
	"clothing-pair-project/config"
	"clothing-pair-project/handler"
	"clothing-pair-project/service"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect db", err.Error())
	}
	defer db.Close()

	for {
		dashboardMenu()

		var input int
		fmt.Print("Choose option: ")
		fmt.Scanln(&input)

		switch input {
		case 1:
			loginMenu(db)
		case 2:
			// TODO: register
		case 0:
			return
		default:
			fmt.Println("Invalid input")
		}
		fmt.Println()
	}
}

func dashboardMenu() {
	fmt.Println()
	fmt.Println("=====================================")
	fmt.Println("Dashboard Menu")
	fmt.Println("1. Login")
	fmt.Println("2. Register")
	fmt.Println("0. Exit")
	fmt.Println("=====================================")
}

func loginMenu(db *sqlx.DB) {
	fmt.Println("=====================================")

	var username, password string
	fmt.Printf("username: ")
	fmt.Scanln(&username)
	fmt.Printf("password: ")
	fmt.Scanln(&password)

	userHandler := handler.NewUserHandler(db)
	userService := service.NewUserService(userHandler)

	user, err := userService.FindByUsername(username)
	if err != nil {
		fmt.Println("Invalid username")
	}
	if user.Password != password {
		fmt.Println("Invalid password")
	} else {
		log.Println("Successfully login")

		fmt.Println()
		fmt.Println("=====================================")
		fmt.Printf("Hi, %s \n", user.Username)
		fmt.Println("=====================================")

		if user.Role == "admin" {
			menu.AdminMenu(db)
		} else {
			// TODO: userMenu(db, user)
		}
	}
}
