package menu

import (
	"bufio"
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
			addUserMenu(userService)
		case 2:
			findAllUsersMenu(userService)
		case 3:
			findUserByUsernameMenu(userService)
		case 4:
			updateUserMenu(userService)
		case 5:
			deleteUserMenu(userService)
		case 0:
			return
		default:
			fmt.Println("Invalid input")
		}
	}
}

func addUserMenu(userService *service.UserService) {
	fmt.Println()
	fmt.Println("=====================================")
	fmt.Println("Add User")
	fmt.Println("=====================================")

	var username, email, password, role string

	for {
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("Username: ")
			username, _ = reader.ReadString('\n')
			username = strings.TrimSpace(username)
			if strings.Contains(username, " ") {
				fmt.Println("Username cannot contain spaces")
				continue
			}
			break
		}

		_, err := userService.FindByUsername(username)
		if err == nil {
			fmt.Println("Username already exists. Please choose another username.")
			continue
		}
		break
	}

	for {
		fmt.Print("Email: ")
		fmt.Scanln(&email)
		if !strings.Contains(email, "@") {
			fmt.Println("Invalid email format. Email must contain @")
			continue
		}
		break
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Password: ")
	password, _ = reader.ReadString('\n')
	password = strings.TrimSpace(password)

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

func findAllUsersMenu(userService *service.UserService) {
	fmt.Println()
	fmt.Println("=====================================")
	fmt.Println("Find All Users")
	fmt.Println("=====================================")

	allUser(userService)

	var input string
	fmt.Print("0. Back: ")
	fmt.Scanln(&input)
	if input == "0" {
		return
	} else {
		fmt.Println("Invalid input")
	}
}

func findUserByUsernameMenu(userService *service.UserService) {
	fmt.Println()
	fmt.Println("=====================================")
	fmt.Println("Find User By Username")
	fmt.Println("=====================================")

	var username string
	fmt.Print("Username: ")
	fmt.Scanln(&username)

	user, err := userService.FindByUsername(username)
	if err != nil {
		fmt.Printf("username %v not found\n", username)
		return
	}

	fmt.Println("User found:")
	fmt.Printf("ID: %d\n", user.UserID)
	fmt.Printf("Username: %s\n", user.Username)
	fmt.Printf("Email: %s\n", user.Email)
	fmt.Printf("Role: %s\n", user.Role)
	fmt.Printf("Created At: %s\n", user.CreatedAt.Format("2006-01-02 15:04:05"))
}

func updateUserMenu(userService *service.UserService) {
	fmt.Println()
	fmt.Println("=====================================")
	fmt.Println("Update User")
	fmt.Println("=====================================")

	var username, email, password, role string

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Username: ")
		username, _ = reader.ReadString('\n')
		username = strings.TrimSpace(username)
		if strings.Contains(username, " ") {
			fmt.Println("Username cannot contain spaces")
			continue
		}
		break
	}

	user, err := userService.FindByUsername(username)
	if err != nil {
		fmt.Printf("username %v not found\n", username)
		return
	}

	for {
		fmt.Print("Email: ")
		fmt.Scanln(&email)
		if !strings.Contains(email, "@") {
			fmt.Println("Invalid email format. Email must contain @")
			continue
		}
		break
	}

	reader = bufio.NewReader(os.Stdin)
	fmt.Print("Password: ")
	password, _ = reader.ReadString('\n')
	password = strings.TrimSpace(password)

	fmt.Print("Role: ")
	fmt.Scanln(&role)

	user.Email = email
	user.Password = password
	user.Role = role

	err = userService.Update(user)
	if err != nil {
		fmt.Printf("Failed to update user: %v\n", err)
		return
	}
	fmt.Println("User updated successfully!")
}

func allUser(userService *service.UserService) {
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
}

func deleteUserMenu(userService *service.UserService) {
	fmt.Println()
	fmt.Println("=====================================")
	fmt.Println("Delete User")
	fmt.Println("=====================================")

	allUser(userService)

	var input string
	fmt.Println("Input ID to delete (0 to back): ")
	fmt.Print("Delete ID: ")
	fmt.Scanln(&input)

	switch input {
	case "0":
		return
	default:
		userID, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}

		err = userService.Delete(userID)
		if err != nil {
			fmt.Printf("Failed to delete user: %v\n", err)
			return
		}
		fmt.Println("User deleted successfully!")
	}
}

func RegisterUser(db *sqlx.DB) {
	fmt.Println()
	fmt.Println("=====================================")
	fmt.Println("Register User")
	fmt.Println("=====================================")

	userHandler := handler.NewUserHandler(db)
	userService := service.NewUserService(userHandler)

	var username, email, password string

	for {
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("Username: ")
			username, _ = reader.ReadString('\n')
			username = strings.TrimSpace(username)
			if strings.Contains(username, " ") {
				fmt.Println("Username cannot contain spaces")
				continue
			}
			break
		}

		_, err := userService.FindByUsername(username)
		if err == nil {
			fmt.Println("Username already exists. Please choose another username.")
			continue
		}
		break
	}

	for {
		fmt.Print("Email: ")
		fmt.Scanln(&email)
		if !strings.Contains(email, "@") {
			fmt.Println("Invalid email format. Email must contain @")
			continue
		}
		break
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Password: ")
	password, _ = reader.ReadString('\n')
	password = strings.TrimSpace(password)

	user := entity.User{
		Username: username,
		Email:    email,
		Password: password,
		Role:     "user",
	}

	err := userService.Add(user)
	if err != nil {
		fmt.Printf("Failed to add user: %v\n", err)
		return
	}
	fmt.Println("User added successfully!")
}
