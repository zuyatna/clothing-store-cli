package menus

import (
	"clothing-pair-project/entity"
	"clothing-pair-project/handler"
	"clothing-pair-project/service"
	"fmt"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/olekukonko/tablewriter"
)

func AdminMenu(db *sqlx.DB) {
	for {
		fmt.Println()
		fmt.Println("=====================================")
		fmt.Println("Admin Menu")
		fmt.Println("1. Manage User")
		fmt.Println("2. Manage Color")
		fmt.Println("3. Manage Payment Method")
		fmt.Println("4. Manage Collection")
		fmt.Println("0. Logout")
		fmt.Println("=====================================")

		var input int
		fmt.Print("Choose option: ")
		fmt.Scanln(&input)

		switch input {
		case 1:
			ManageUserMenu(db)
		case 2:
			// TODO: manageColorMenu(db)
		case 3:
			// TODO: managePaymentMethodMenu(db)
		case 4:
			// TODO: manageCollectionMenu(db)
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
	fmt.Print("Email: ")
	fmt.Scanln(&email)
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
	userService.Add(user)
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
