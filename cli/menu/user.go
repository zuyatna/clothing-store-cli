package menu

import (
	"clothing-pair-project/entity"
	"clothing-pair-project/handler"
	"clothing-pair-project/service"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/olekukonko/tablewriter"
)

func AdminMenu(db *sqlx.DB) {
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
		fmt.Println("0. Logout")
		fmt.Println("=====================================")

		var input int
		fmt.Print("Choose option: ")
		fmt.Scanln(&input)

		switch input {
		case 1:
			ManageUserMenu(db)
		case 2:
			// TODO:
		case 3:
			// TODO:
		case 4:
			// TODO:
		case 0:
			return
		default:
			fmt.Println("Invalid input")
		}
	}
}

func ManageUserMenu(db *sqlx.DB) {
	for {
		fmt.Println()
		fmt.Println("=====================================")
		fmt.Println("Manage User Menu")
		fmt.Println("1. Add User")
		fmt.Println("2. Find All Users")
		fmt.Println("3. Find User By Username")
		fmt.Println("4. Update User")
		fmt.Println("5. Delete User")
		fmt.Println("0. Back")
		fmt.Println("=====================================")

		var input int
		fmt.Print("Choose option: ")
		fmt.Scanln(&input)

		userHandler := handler.NewUserHandler(db)
		userService := service.NewUserService(userHandler)

		switch input {
		case 1:
			AddUserMenu(userService)
		case 2:
			FindAllUsersMenu(userService)
		case 3:
			// TODO: findUserByUsernameMenu(userService)
		case 4:
			// TODO: updateUserMenu(userService)
		case 5:
			// TODO: deleteUserMenu(userService)
		case 0:
			return
		default:
			fmt.Println("Invalid input")
		}
	}
}

func AddUserMenu(userService *service.UserService) {
	fmt.Println()
	fmt.Println("=====================================")
	fmt.Println("Add User")
	fmt.Println("=====================================")

	var username, email, password, role string
	fmt.Print("Username: ")
	fmt.Scanln(&username)

	for {
		fmt.Print("Email: ")
		fmt.Scanln(&email)
		if !strings.Contains(email, "@") {
			fmt.Println("Invalid email format. Email must contain @")
			continue
		}
		break
	}

	fmt.Print("Password: ")
	fmt.Scanln(&password)
	fmt.Print("Role: ")
	fmt.Scanln(&role)

	user := entity.User{
		Username: username,
		Email:    email,
		Password: password,
		Role:     role,
	}

	err := userService.Add(user)
	if err != nil {
		fmt.Printf("Failed to add user: %v\n", err)
		return
	}
	fmt.Println("User added successfully!")
}

func FindAllUsersMenu(userService *service.UserService) {
	fmt.Println()
	fmt.Println("=====================================")
	fmt.Println("Find All Users")
	fmt.Println("=====================================")

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Username", "Email", "Role", "Created At"})

	users, err := userService.FindAll()
	if err != nil {
		fmt.Println("Error fetching users:", err)
		return
	}

	if len(users) == 0 {
		fmt.Println("No users found.")
		return
	}

	for _, user := range users {
		table.Append([]string{
			strconv.Itoa(user.UserID),
			user.Username,
			user.Email,
			user.Role,
			user.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	table.Render()

	fmt.Println()

	var input string
	fmt.Print("0. Back: ")
	fmt.Scanln(&input)
	if input == "0" {
		return
	} else {
		fmt.Println("Invalid input")
	}
}
