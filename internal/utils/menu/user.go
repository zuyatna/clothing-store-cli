package menu

import (
	"bufio"
	"clothing-pair-project/internal/database/sql"
	"clothing-pair-project/internal/helper"
	"clothing-pair-project/internal/models"
	"clothing-pair-project/internal/services"
	"clothing-pair-project/internal/utils/messages"
	"clothing-pair-project/internal/utils/terminal"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/olekukonko/tablewriter"

	_ "github.com/lib/pq"
)

func ManageUserMenu(db *sqlx.DB, message string) {
	terminal.Clear()

	fmt.Println()
	fmt.Println("=====================================")
	fmt.Println("Manage User Menu")
	fmt.Println("1. Find All Users")
	fmt.Println("2. Find User By Username")
	fmt.Println("3. Add User")
	fmt.Println("4. Update User")
	fmt.Println("5. Delete User")
	fmt.Println("0. Back")
	fmt.Println("=====================================")

	messages.PrintMessage(message)

	var input string
	fmt.Print("Choose option: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		message = "No input entered"
		messages.PrintMessage(message)
		ManageUserMenu(db, message)
	}

	userRepository := sql.NewUserRepository(db)
	userService := services.NewUserService(userRepository)

	switch input {
	case "1":
		findAllUsersMenu(db, userService, "")
	case "2":
		findUserByUsername(db, userService, "")
	case "3":
		addUserMenu(db, userService, "")
	case "4":
		updateUserMenu(db, userService, "")
	case "5":
		deleteUserMenu(db, userService, "")
	case "0":
		AdminMenu(db, "")
	default:
		message = "Invalid input"
		messages.PrintMessage(message)
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

func findAllUsersMenu(db *sqlx.DB, userService *services.UserService, message string) {
	terminal.Clear()

	fmt.Println("=====================================")
	fmt.Println("Find All Users")
	fmt.Println("=====================================")

	allUser(userService)

	messages.PrintMessage(message)

	var input string
	fmt.Print("0. Back: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		message = "No input entered"
		messages.PrintMessage(message)
		findAllUsersMenu(db, userService, message)
	}

	if input == "0" {
		ManageUserMenu(db, "")
	} else {
		message = "Invalid input"
		messages.PrintMessage(message)
		findAllUsersMenu(db, userService, message)
	}
}

func findUserByUsername(db *sqlx.DB, userService *services.UserService, message string) {
	terminal.Clear()

	fmt.Println("=====================================")
	fmt.Println("Find User By Username")
	fmt.Println("=====================================")

	reader := bufio.NewReader(os.Stdin)

	messages.PrintMessage(message)

	fmt.Print("Enter username: ")
	username, err := reader.ReadString('\n')
	username = strings.TrimSpace(username)
	if err != nil {
		message = "Error reading input"
		messages.PrintMessage(message)
		findUserByUsername(db, userService, message)
	}

	if username == "" {
		message = "Username cannot be empty"
		messages.PrintMessage(message)
		findUserByUsername(db, userService, message)
	} else if strings.Contains(username, " ") {
		message = "Username cannot contain spaces"
		messages.PrintMessage(message)
		findUserByUsername(db, userService, message)
	}
	fmt.Println()

	user, err := userService.GetUserByUsername(username)
	if err != nil {
		message = "User not found"
		messages.PrintMessage(message)
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
		messages.PrintMessage(message)
		ManageUserMenu(db, message)
	}
	ManageUserMenu(db, "")
}

func addUserMenu(db *sqlx.DB, userService *services.UserService, message string) {
	terminal.Clear()

	userRepository := sql.NewUserRepository(db)

	fmt.Println("=====================================")
	fmt.Println("Add User")
	fmt.Println("=====================================")

	messages.PrintMessage(message)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter username: ")
	username, err := reader.ReadString('\n')
	username = strings.TrimSpace(username)
	if err != nil {
		message = "Error reading enumTrim"
		messages.PrintMessage(message)
		addUserMenu(db, userService, message)
	}

	if username == "" {
		message = "Username cannot be empty"
		messages.PrintMessage(message)
		addUserMenu(db, userService, message)
	} else if strings.Contains(username, " ") {
		message = "Username cannot contain spaces"
		messages.PrintMessage(message)
		addUserMenu(db, userService, message)
	}

	fmt.Print("Enter email: ")
	email, err := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	if err != nil {
		message = "Error reading enumTrim"
		messages.PrintMessage(message)
		addUserMenu(db, userService, message)
	}

	if email == "" {
		message = "Email cannot be empty"
		messages.PrintMessage(message)
		addUserMenu(db, userService, message)
	} else if strings.Contains(email, " ") {
		message = "Email cannot contain spaces"
		messages.PrintMessage(message)
		addUserMenu(db, userService, message)
	} else if !strings.Contains(email, "@") {
		message = "Invalid email format"
		messages.PrintMessage(message)
		addUserMenu(db, userService, message)
	}

	password, err := terminal.HidePassword("Enter Password: ")
	if err != nil {
		message = "Error reading password"
		messages.PrintMessage(message)
		addUserMenu(db, userService, message)
	}

	if password == nil {
		message = "Password cannot be empty"
		messages.PrintMessage(message)
		addUserMenu(db, userService, message)
	} else if strings.Contains(string(password), " ") {
		message = "Password cannot contain spaces"
		messages.PrintMessage(message)
		addUserMenu(db, userService, message)
	}

	confirmPassword, err := terminal.HidePassword("Confirm Password:")
	if err != nil {
		message = "Error reading password"
		messages.PrintMessage(message)
		addUserMenu(db, userService, message)
	}

	if string(password) != string(confirmPassword) {
		message = "Password and confirm password do not match"
		messages.PrintMessage(message)
		addUserMenu(db, userService, message)
	}

	enumRange, err := userRepository.EnumRole()
	if err != nil {
		message = "Error fetching enumRange"
		messages.PrintMessage(message)
		addUserMenu(db, userService, message)
	}

	enumTrim := enumRange
	enumTrim = strings.Trim(enumTrim, "{}")
	enumStr := strings.ReplaceAll(enumTrim, ",", " ")
	roles := strings.Split(enumStr, " ")

	fmt.Printf("Role %s: ", enumRange)
	role, err := reader.ReadString('\n')
	role = strings.TrimSpace(role)
	if err != nil {
		message = "Error reading enumTrim"
		messages.PrintMessage(message)
		addUserMenu(db, userService, message)
	}

	if role == "" {
		message = "Role cannot be empty"
		messages.PrintMessage(message)
		addUserMenu(db, userService, message)
	}

	if !slices.Contains(roles, role) {
		message = "Role not found"
		messages.PrintMessage(message)
		addUserMenu(db, userService, message)
	}

	fmt.Print("Confirm Add User? (y/n): ")
	confirm, err := reader.ReadString('\n')
	confirm = strings.TrimSpace(confirm)
	if err != nil {
		message = "Error reading enumTrim"
		messages.PrintMessage(message)
		addUserMenu(db, userService, message)
	}

	if confirm != "y" {
		ManageUserMenu(db, "Add user cancelled, returning to Manage User Menu")
	}
	fmt.Println()

	hashedPassword, err := helper.HashPassword(string(password))
	if err != nil {
		message = "Error hashing password"
		messages.PrintMessage(message)
		addUserMenu(db, userService, message)
	}

	user := models.User{
		Username: username,
		Email:    email,
		Password: hashedPassword,
		Role:     role,
	}

	err = userService.AddUser(user)
	if err != nil {
		if strings.Contains(err.Error(), "users_email_key") {
			message = "Email already exists"
		} else if strings.Contains(err.Error(), "users_username_key") {
			message = "Username already exists"
		} else {
			message = fmt.Sprintf("Error adding user: %v", err)
		}
		messages.PrintMessage(message)
		addUserMenu(db, userService, message)
	}

	ManageUserMenu(db, "Successfully added user with username "+username)
}

func updateUserMenu(db *sqlx.DB, userService *services.UserService, message string) {
	terminal.Clear()

	fmt.Println("=====================================")
	fmt.Println("Update User")
	fmt.Println("=====================================")

	allUser(userService)

	messages.PrintMessage(message)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter User ID: ")
	userID, err := reader.ReadString('\n')
	if err != nil {
		message = "Error reading input"
		messages.PrintMessage(message)
		updateUserMenu(db, userService, message)
	}

	if userID == "" {
		message = "User ID cannot be empty"
		messages.PrintMessage(message)
		updateUserMenu(db, userService, message)
	} else if strings.Contains(userID, " ") {
		message = "User ID cannot contain spaces"
		messages.PrintMessage(message)
		updateUserMenu(db, userService, message)
	}

	userIDInt, err := strconv.Atoi(strings.TrimSpace(userID))
	if err != nil {
		message = "User ID must be a number"
		messages.PrintMessage(message)
		updateUserMenu(db, userService, message)
	}

	user, err := userService.GetUserByID(userIDInt)
	if err != nil {
		message = "User not found"
		messages.PrintMessage(message)
		updateUserMenu(db, userService, message)
	}

	fmt.Println("Current Username:", user.Username)

	fmt.Print("Do you want to update username? (y/n): ")
	updateUsername, err := reader.ReadString('\n')
	updateUsername = strings.TrimSpace(updateUsername)
	if err != nil {
		message = "Error reading input"
		messages.PrintMessage(message)
		updateUserMenu(db, userService, message)
	}

	if updateUsername == "y" {
		fmt.Print("Enter new username: ")
		username, err := reader.ReadString('\n')
		username = strings.TrimSpace(username)
		if err != nil {
			message = "Error reading input"
			messages.PrintMessage(message)
			updateUserMenu(db, userService, message)
		}

		if username == "" {
			message = "Username cannot be empty"
			messages.PrintMessage(message)
			updateUserMenu(db, userService, message)
		} else if strings.Contains(username, " ") {
			message = "Username cannot contain spaces"
			messages.PrintMessage(message)
			updateUserMenu(db, userService, message)
		}

		user.Username = username
	} else if updateUsername != "n" {
		message = "Invalid input"
		messages.PrintMessage(message)
		updateUserMenu(db, userService, message)
	}

	fmt.Println("Current Email:", user.Email)

	fmt.Print("Do you want to update email? (y/n): ")
	updateEmail, err := reader.ReadString('\n')
	updateEmail = strings.TrimSpace(updateEmail)
	if err != nil {
		message = "Error reading input"
		messages.PrintMessage(message)
		updateUserMenu(db, userService, message)
	}

	if updateEmail == "y" {
		fmt.Print("Enter new email: ")
		email, err := reader.ReadString('\n')
		email = strings.TrimSpace(email)
		if err != nil {
			message = "Error reading input"
			messages.PrintMessage(message)
			updateUserMenu(db, userService, message)
		}

		if email == "" {
			message = "Email cannot be empty"
			messages.PrintMessage(message)
			updateUserMenu(db, userService, message)
		} else if strings.Contains(email, " ") {
			message = "Email cannot contain spaces"
			messages.PrintMessage(message)
			updateUserMenu(db, userService, message)
		} else if !strings.Contains(email, "@") {
			message = "Invalid email format"
			messages.PrintMessage(message)
			updateUserMenu(db, userService, message)
		}

		user.Email = email
	} else if updateEmail != "n" {
		message = "Invalid input"
		messages.PrintMessage(message)
		updateUserMenu(db, userService, message)
	}

	fmt.Print("Do you want to update password? (y/n): ")
	updatePassword, err := reader.ReadString('\n')
	updatePassword = strings.TrimSpace(updatePassword)
	if err != nil {
		message = "Error reading input"
		messages.PrintMessage(message)
		updateUserMenu(db, userService, message)
	}

	if updatePassword == "y" {
		password, err := terminal.HidePassword("Enter new password (min 8 characters): ")
		if err != nil {
			message = "Error reading password"
			messages.PrintMessage(message)
			updateUserMenu(db, userService, message)
		}

		if password == nil {
			message = "Password cannot be empty"
			messages.PrintMessage(message)
			updateUserMenu(db, userService, message)
		} else if len(password) < 8 {
			message = "Password must be at least 8 characters"
			messages.PrintMessage(message)
			updateUserMenu(db, userService, message)
		} else if strings.Contains(string(password), " ") {
			message = "Password cannot contain spaces"
			messages.PrintMessage(message)
			updateUserMenu(db, userService, message)
		}

		confirmPassword, err := terminal.HidePassword("Confirm new password (min 8 characters):")
		if err != nil {
			message = "Error reading password"
			messages.PrintMessage(message)
			updateUserMenu(db, userService, message)
		}

		if string(password) != string(confirmPassword) {
			message = "Password and confirm password do not match"
			messages.PrintMessage(message)
			updateUserMenu(db, userService, message)
		}

		user.Password = string(password)
	} else if updatePassword != "n" {
		message = "Invalid input"
		messages.PrintMessage(message)
		updateUserMenu(db, userService, message)
	}

	err = userService.UpdateUser(user)
	if err != nil {
		if strings.Contains(err.Error(), "users_email_key") {
			message = "Email already exists"
		} else if strings.Contains(err.Error(), "users_username_key") {
			message = "Username already exists"
		} else {
			message = "Error updating user"
		}
		messages.PrintMessage(message)
		updateUserMenu(db, userService, message)
	}

	ManageUserMenu(db, "Successfully updated user with ID "+strconv.Itoa(user.UserID))
}

func deleteUserMenu(db *sqlx.DB, userService *services.UserService, message string) {
	terminal.Clear()

	fmt.Println("=====================================")
	fmt.Println("Delete User")
	fmt.Println("=====================================")

	messages.PrintMessage(message)

	allUser(userService)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println()
	fmt.Print("Enter User ID: ")
	userID, err := reader.ReadString('\n')
	if err != nil {
		message = "Error reading input"
		messages.PrintMessage(message)
		deleteUserMenu(db, userService, message)
	}

	if userID == "" {
		message = "User ID cannot be empty"
		messages.PrintMessage(message)
		deleteUserMenu(db, userService, message)
	} else if strings.Contains(userID, " ") {
		message = "User ID cannot contain spaces"
		messages.PrintMessage(message)
		deleteUserMenu(db, userService, message)
	}

	userIDInt, err := strconv.Atoi(strings.TrimSpace(userID))
	if err != nil {
		message = "Invalid User ID"
		messages.PrintMessage(message)
		deleteUserMenu(db, userService, message)
	}

	err = userService.DeleteUser(userIDInt)
	if err != nil {
		message = "Error deleting user"
		messages.PrintMessage(message)
		deleteUserMenu(db, userService, message)
	}

	ManageUserMenu(db, "Successfully deleted user with ID "+userID)
}
