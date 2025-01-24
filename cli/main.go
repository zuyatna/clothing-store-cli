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
			menu.RegisterUser(db)
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
	} else if user.Password != password {
		fmt.Println("Invalid password")
	} else {
		log.Println("Successfully login")

		fmt.Println()
		fmt.Println("=====================================")
		fmt.Printf("Hi, %s \n", user.Username)
		fmt.Println("=====================================")

		if user.Role == "admin" {
			adminMenu(db)
		} else {
			// TODO: userMenu(db, user)
		}
	}
}

func adminMenu(db *sqlx.DB) {
	for {
		fmt.Println()
		fmt.Println("=====================================")
		fmt.Println("Admin Menu")
		fmt.Println("1. Manage User")
		fmt.Println("2. Manage Product")
		fmt.Println("3. Manage Collection")
		fmt.Println("4. Manage Category")
		fmt.Println("5. Manage Color")
		fmt.Println("6. Manage Size")
		fmt.Println("7. Manage Payment Method")
		fmt.Println("8. Reports")
		fmt.Println("0. Logout")
		fmt.Println("=====================================")

		var input int
		fmt.Print("Choose option: ")
		fmt.Scanln(&input)

		switch input {
		case 1:
			menu.ManageUserMenu(db)
		case 2:
			menu.ManageProductMenu(db)
		case 3:
			menu.ManageCollectionMenu(db)
		case 4:
			menu.ManageCategoryMenu(db)
		case 5:
			menu.ManageColorMenu(db)
		case 6:
			// TODO: manageSizeMenu(db)
		case 7:
			// TODO: managePaymentMethodMenu(db)
		case 8:
			menu.ReportMenu(db)
		case 0:
			return
		default:
			fmt.Println("Invalid input")
		}
	}
}
