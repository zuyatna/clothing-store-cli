package menu

import (
	"bufio"
	"clothing-pair-project/internal/database/sql"
	"clothing-pair-project/internal/models"
	"clothing-pair-project/internal/services"
	"clothing-pair-project/internal/utils/terminal"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/olekukonko/tablewriter"
	"os"
	"slices"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

func ManageUserMenu(db *sqlx.DB, message string) {
	terminal.Clear()

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

	fmt.Println()
	fmt.Println(message)
	fmt.Println()

	var input string
	fmt.Print("Choose option: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		message = "No input entered"
		fmt.Println(message)
		fmt.Println()

		ManageUserMenu(db, message)
	}

	userRepository := sql.NewUserRepository(db)
	userService := services.NewUserService(userRepository)

	switch input {
	case "1":
		addUserMenu(db, userService, "")
	case "2":
		findAllUsersMenu(db, userService, "")
	case "3":
		findUserByUsername(db, userService, "")
	case "4":
		// TODO: update user menu
	case "5":
		deleteUserMenu(db, userService, "")
	case "0":
		AdminMenu(db, "")
	default:
		message = "Invalid input"
		fmt.Println(message)
		fmt.Println()

		ManageUserMenu(db, message)
	}

}

func allUser(userService *services.UserService) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Username", "Email", "Role", "Created At", "Active"})

	users, err := userService.GetAllUsers()
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
			strconv.FormatBool(user.Active),
		})
	}
	table.Render()

	fmt.Println()
}

func addUserMenu(db *sqlx.DB, userService *services.UserService, message string) {
	terminal.Clear()

	userRepository := sql.NewUserRepository(db)

	fmt.Println("=====================================")
	fmt.Println("Add User")
	fmt.Println("=====================================")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println()
	fmt.Println(message)
	fmt.Println()

	fmt.Print("Enter username: ")
	username, err := reader.ReadString('\n')
	username = strings.TrimSpace(username)
	if err != nil {
		message = "Error reading input"
		fmt.Println(message)
		fmt.Println()

		addUserMenu(db, userService, message)
	}

	if username == "" {
		message = "Username cannot be empty"
		fmt.Println(message)
		fmt.Println()

		addUserMenu(db, userService, message)
	} else if strings.Contains(username, " ") {
		message = "Username cannot contain spaces"
		fmt.Println(message)
		fmt.Println()

		addUserMenu(db, userService, message)
	}

	user, _ := userService.GetUserByUsername(username)

	if user.Username == username {
		message = "Username already exists"
		fmt.Println(message)
		fmt.Println()

		addUserMenu(db, userService, message)
	}

	fmt.Print("Enter email: ")
	email, err := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	if err != nil {
		message = "Error reading input"
		fmt.Println(message)
		fmt.Println()

		addUserMenu(db, userService, message)
	}

	if email == "" {
		message = "Email cannot be empty"
		fmt.Println(message)
		fmt.Println()

		addUserMenu(db, userService, message)
	} else if strings.Contains(email, " ") {
		message = "Email cannot contain spaces"
		fmt.Println(message)
		fmt.Println()

		addUserMenu(db, userService, message)
	} else if !strings.Contains(email, "@") {
		message = "Invalid email format"
		fmt.Println(message)
		fmt.Println()

		addUserMenu(db, userService, message)
	}

	password, err := terminal.HidePassword("Enter Password (min 8 characters): ")
	if err != nil {
		message = "Error reading password"
		fmt.Println(message)
		fmt.Println()

		addUserMenu(db, userService, message)
	}

	if password == nil {
		message = "Password cannot be empty"
		fmt.Println(message)
		fmt.Println()

		addUserMenu(db, userService, message)
	} else if len(password) < 8 {
		message = "Password must be at least 8 characters"
		fmt.Println(message)
		fmt.Println()

		addUserMenu(db, userService, message)
	} else if strings.Contains(string(password), " ") {
		message = "Password cannot contain spaces"
		fmt.Println(message)
		fmt.Println()

		addUserMenu(db, userService, message)
	}

	confirmPassword, err := terminal.HidePassword("Confirm Password (min 8 characters):")
	if err != nil {
		message = "Error reading password"
		fmt.Println(message)
		fmt.Println()

		addUserMenu(db, userService, message)
	}

	if string(password) != string(confirmPassword) {
		message = "Password and confirm password do not match"
		fmt.Println(message)
		fmt.Println()

		addUserMenu(db, userService, message)
	}

	enumRange, err := userRepository.EnumRole()
	if err != nil {
		message = "Error fetching enumRange"
		fmt.Println(message)
		fmt.Println()

		addUserMenu(db, userService, message)
	}

	input := enumRange
	input = strings.Trim(input, "{}")
	output := strings.ReplaceAll(input, ",", " ")
	roles := strings.Split(output, " ")

	fmt.Printf("Role %s: ", enumRange)
	role, err := reader.ReadString('\n')
	role = strings.TrimSpace(role)
	if err != nil {
		message = "Error reading input"
		fmt.Println(message)
		fmt.Println()

		addUserMenu(db, userService, message)
	}

	if role == "" {
		message = "Role cannot be empty"
		fmt.Println(message)
		fmt.Println()

		addUserMenu(db, userService, message)
	}

	if !slices.Contains(roles, role) {
		message = "Role not found"
		fmt.Println(message)
		fmt.Println()

		addUserMenu(db, userService, message)
	}

	fmt.Print("Confirm Add User? (y/n): ")
	confirm, err := reader.ReadString('\n')
	confirm = strings.TrimSpace(confirm)
	if err != nil {
		message = "Error reading input"
		fmt.Println(message)
		fmt.Println()

		addUserMenu(db, userService, message)
	}

	if confirm != "y" {
		ManageUserMenu(db, "Add user cancelled, returning to Manage User Menu")
	}
	fmt.Println()

	user = models.User{
		Username: username,
		Email:    email,
		Password: string(password),
		Role:     role,
		Active:   true,
	}

	err = userService.AddUser(user)
	if err != nil {
		if strings.Contains(err.Error(), "users_email_key") {
			message = "Email already exists"
		} else {
			message = "Error adding user"
		}

		fmt.Println(message)
		fmt.Println()

		addUserMenu(db, userService, message)
	}

	ManageUserMenu(db, "Successfully added user with username "+username)
}

func findAllUsersMenu(db *sqlx.DB, userService *services.UserService, message string) {
	terminal.Clear()

	fmt.Println("=====================================")
	fmt.Println("Find All Users")
	fmt.Println("=====================================")

	allUser(userService)

	fmt.Println()
	fmt.Println(message)
	fmt.Println()

	var input string
	fmt.Print("0. Back: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		message = "No input entered"
		fmt.Println(message)
		fmt.Println()

		findAllUsersMenu(db, userService, message)
	}

	if input == "0" {
		ManageUserMenu(db, "")
	} else {
		message = "Invalid input"
		fmt.Println(message)
		fmt.Println()

		findAllUsersMenu(db, userService, message)
	}
}

func findUserByUsername(db *sqlx.DB, userService *services.UserService, message string) {
	terminal.Clear()

	fmt.Println("=====================================")
	fmt.Println("Find User By Username")
	fmt.Println("=====================================")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println()
	fmt.Println(message)
	fmt.Println()

	fmt.Print("Enter username: ")
	username, err := reader.ReadString('\n')
	username = strings.TrimSpace(username)
	if err != nil {
		message = "Error reading input"
		fmt.Println(message)
		fmt.Println()

		findUserByUsername(db, userService, message)
	}

	if username == "" {
		message = "Username cannot be empty"
		fmt.Println(message)
		fmt.Println()

		findUserByUsername(db, userService, message)
	} else if strings.Contains(username, " ") {
		message = "Username cannot contain spaces"
		fmt.Println(message)
		fmt.Println()

		findUserByUsername(db, userService, message)
	}
	fmt.Println()

	user, err := userService.GetUserByUsername(username)
	if err != nil {
		message = "User not found"
		fmt.Println(message)
		fmt.Println()

		findUserByUsername(db, userService, message)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Username", "Email", "Role", "Created At", "Active"})
	table.Append([]string{
		strconv.Itoa(user.UserID),
		user.Username,
		user.Email,
		user.Role,
		user.CreatedAt.Format("2006-01-02 15:04:05"),
		strconv.FormatBool(user.Active),
	})
	table.Render()

	fmt.Println()
	fmt.Print("Press any key to back... ")
	_, err = bufio.NewReader(os.Stdin).ReadBytes('\n')
	if err != nil {
		message = "Error reading input"
		fmt.Println(message)
		fmt.Println()

		ManageUserMenu(db, message)
	}
	ManageUserMenu(db, "")
}

func deleteUserMenu(db *sqlx.DB, userService *services.UserService, message string) {
	terminal.Clear()

	fmt.Println("=====================================")
	fmt.Println("Delete User")
	fmt.Println("=====================================")

	fmt.Println()
	fmt.Println(message)
	fmt.Println()

	allUser(userService)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println()
	fmt.Print("Enter User ID: ")
	userID, err := reader.ReadString('\n')
	if err != nil {
		message = "Error reading input"
		fmt.Println(message)
		fmt.Println()

		deleteUserMenu(db, userService, message)
	}

	if userID == "" {
		message = "User ID cannot be empty"
		fmt.Println(message)
		fmt.Println()

		deleteUserMenu(db, userService, message)
	} else if strings.Contains(userID, " ") {
		message = "User ID cannot contain spaces"
		fmt.Println(message)
		fmt.Println()

		deleteUserMenu(db, userService, message)
	}

	userIDInt, err := strconv.Atoi(strings.TrimSpace(userID))
	if err != nil {
		message = "Invalid User ID"
		fmt.Println(message)
		fmt.Println()

		deleteUserMenu(db, userService, message)
	}

	err = userService.DeleteUser(userIDInt)
	if err != nil {
		message = "Error deleting user"
		fmt.Println(message)
		fmt.Println()

		deleteUserMenu(db, userService, message)
	}

	ManageUserMenu(db, "Successfully deleted user with ID "+userID)
}
