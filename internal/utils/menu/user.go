package menu

import (
	"bufio"
	"clothing-pair-project/internal/database/sqlrepo"
	"clothing-pair-project/internal/models"
	"clothing-pair-project/internal/services"
	"clothing-pair-project/internal/utils/handler"
	"clothing-pair-project/internal/utils/input"
	"clothing-pair-project/internal/utils/messages"
	"clothing-pair-project/internal/utils/tables"
	"clothing-pair-project/internal/utils/terminal"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/olekukonko/tablewriter"

	_ "github.com/lib/pq"
)

func ManageUserMenu(db *sqlx.DB, msg string) {
	fmt.Println()
	fmt.Println("=====================================")
	fmt.Println("Manage User Menu")
	fmt.Println("1. Find All Users")
	fmt.Println("2. Find User By Username")
	fmt.Println("3. Add User")
	fmt.Println("4. Edit User")
	fmt.Println("5. Delete User")
	fmt.Println("0. Back")
	fmt.Println("=====================================")

	messages.PrintMessage(msg)

	userRepository := sqlrepo.NewUserRepository(db)
	userService := services.NewUserService(userRepository)

	var input string
	fmt.Print("Choose option: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		msg = "No input entered"
		ManageUserMenu(db, msg)
	}

	switch input {
	case "1":
		findAllUsersMenu(db, userService)
	case "2":
		findUserByUsername(db, userService, "")
	case "3":
		addUserMenu(db, userService, "")
	case "4":
		editUserMenu(db, userService, "")
	case "5":
		deleteUserMenu(db, userService, "")
	case "0":
		AdminMenu(db, "")
	default:
		msg = "Invalid input"
		ManageUserMenu(db, msg)
	}
}

func showUsers(userService *services.UserService) {
	writer := tablewriter.NewWriter(os.Stdout)
	displayer := tables.NewTableUsersDisplayer(writer)
	handler := handler.NewUserHandler(userService, displayer)
	if err := handler.ShowAllUsers(); err != nil {
		fmt.Printf("Error fetching all users: %v\n", err)
	}
}

func findAllUsersMenu(db *sqlx.DB, userService *services.UserService) {
	fmt.Println("=====================================")
	fmt.Println("Find All Users")
	fmt.Println("=====================================")

	showUsers(userService)

	input.BackMenu()
	ManageUserMenu(db, "")
}

func findUserByUsername(db *sqlx.DB, userService *services.UserService, msg string) {
	fmt.Println("=====================================")
	fmt.Println("Find User By Username")
	fmt.Println("=====================================")

	messages.PrintMessage(msg)

	username, err := input.Username()
	if err != nil {
		msg = err.Error()
		findUserByUsername(db, userService, msg)
	}
	fmt.Println()

	writer := tablewriter.NewWriter(os.Stdout)
	displayer := tables.NewTableUsersDisplayer(writer)
	handler := handler.NewUserHandler(userService, displayer)
	if err := handler.ShowUserByUsername(username); err != nil {
		msg = "Error fetching user by username"
		findUserByUsername(db, userService, msg)
	}

	input.BackMenu()
	ManageUserMenu(db, "")
}

func addUserMenu(db *sqlx.DB, userService *services.UserService, msg string) {
	fmt.Println("=====================================")
	fmt.Println("Add User")
	fmt.Println("=====================================")

	messages.PrintMessage(msg)

	username, err := input.Username()
	if err != nil {
		msg = err.Error()
		addUserMenu(db, userService, msg)
	}

	email, err := input.Email()
	if err != nil {
		msg = err.Error()
		addUserMenu(db, userService, msg)
	}

	password, err := input.Password()
	if err != nil {
		msg = err.Error()
		addUserMenu(db, userService, msg)
	}

	role, err := input.Role(db)
	if err != nil {
		msg = err.Error()
		addUserMenu(db, userService, msg)
	}

	user := models.User{
		Username: username,
		Email:    email,
		Password: password,
		Role:     role,
	}

	writer := tablewriter.NewWriter(os.Stdout)
	displayer := tables.NewTableAddUserDisplayer(writer)
	displayer.DisplayAddUser(user)

	confirm, err := input.ConfirmAddUser()
	if err != nil {
		msg = err.Error()
		addUserMenu(db, userService, msg)
	}

	if !confirm {
		msg = "User not added"
		addUserMenu(db, userService, msg)
	}

	err = userService.AddUser(user)
	if err != nil {
		if strings.Contains(err.Error(), "users_email_key") {
			msg = "Email already exists"
		} else if strings.Contains(err.Error(), "users_username_key") {
			msg = "Username already exists"
		} else {
			msg = fmt.Sprintf("Error adding user: %v", err)
		}
		addUserMenu(db, userService, msg)
	}
	msg = "User added successfully"
	ManageUserMenu(db, msg)
}

func editUserMenu(db *sqlx.DB, userService *services.UserService, msg string) {
	fmt.Println("=====================================")
	fmt.Println("Edit User")
	fmt.Println("=====================================")

	showUsers(userService)

	messages.PrintMessage(msg)

	fmt.Print("Enter User ID: ")
	reader := bufio.NewReader(os.Stdin)
	userID, err := reader.ReadString('\n')
	if err != nil {
		msg = "Error reading input"
		messages.PrintMessage(msg)
		editUserMenu(db, userService, msg)
	}

	if userID == "" {
		msg = "User ID cannot be empty"
		messages.PrintMessage(msg)
		editUserMenu(db, userService, msg)
	} else if strings.Contains(userID, " ") {
		msg = "User ID cannot contain spaces"
		messages.PrintMessage(msg)
		editUserMenu(db, userService, msg)
	}

	userIDInt, err := strconv.Atoi(strings.TrimSpace(userID))
	if err != nil {
		msg = "User ID must be a number"
		messages.PrintMessage(msg)
		editUserMenu(db, userService, msg)
	}

	user, err := userService.GetUserByID(userIDInt)
	if err != nil {
		msg = "User not found"
		messages.PrintMessage(msg)
		editUserMenu(db, userService, msg)
	}

	fmt.Println("Current Username:", user.Username)

	fmt.Print("Do you want to update username? (y/n): ")
	updateUsername, err := reader.ReadString('\n')
	updateUsername = strings.TrimSpace(updateUsername)
	if err != nil {
		msg = "Error reading input"
		messages.PrintMessage(msg)
		editUserMenu(db, userService, msg)
	}

	if updateUsername == "y" {
		fmt.Print("Enter new username: ")
		username, err := reader.ReadString('\n')
		username = strings.TrimSpace(username)
		if err != nil {
			msg = "Error reading input"
			messages.PrintMessage(msg)
			editUserMenu(db, userService, msg)
		}

		if username == "" {
			msg = "Username cannot be empty"
			messages.PrintMessage(msg)
			editUserMenu(db, userService, msg)
		} else if strings.Contains(username, " ") {
			msg = "Username cannot contain spaces"
			messages.PrintMessage(msg)
			editUserMenu(db, userService, msg)
		}

		user.Username = username
	} else if updateUsername != "n" {
		msg = "Invalid input"
		messages.PrintMessage(msg)
		editUserMenu(db, userService, msg)
	}

	fmt.Println("Current Email:", user.Email)

	fmt.Print("Do you want to update email? (y/n): ")
	updateEmail, err := reader.ReadString('\n')
	updateEmail = strings.TrimSpace(updateEmail)
	if err != nil {
		msg = "Error reading input"
		messages.PrintMessage(msg)
		editUserMenu(db, userService, msg)
	}

	if updateEmail == "y" {
		fmt.Print("Enter new email: ")
		email, err := reader.ReadString('\n')
		email = strings.TrimSpace(email)
		if err != nil {
			msg = "Error reading input"
			messages.PrintMessage(msg)
			editUserMenu(db, userService, msg)
		}

		if email == "" {
			msg = "Email cannot be empty"
			messages.PrintMessage(msg)
			editUserMenu(db, userService, msg)
		} else if strings.Contains(email, " ") {
			msg = "Email cannot contain spaces"
			messages.PrintMessage(msg)
			editUserMenu(db, userService, msg)
		} else if !strings.Contains(email, "@") {
			msg = "Invalid email format"
			messages.PrintMessage(msg)
			editUserMenu(db, userService, msg)
		}

		user.Email = email
	} else if updateEmail != "n" {
		msg = "Invalid input"
		messages.PrintMessage(msg)
		editUserMenu(db, userService, msg)
	}

	fmt.Print("Do you want to update password? (y/n): ")
	updatePassword, err := reader.ReadString('\n')
	updatePassword = strings.TrimSpace(updatePassword)
	if err != nil {
		msg = "Error reading input"
		messages.PrintMessage(msg)
		editUserMenu(db, userService, msg)
	}

	if updatePassword == "y" {
		password, err := terminal.HidePassword("Enter new password (min 6 characters): ")
		if err != nil {
			msg = "Error reading password"
			messages.PrintMessage(msg)
			editUserMenu(db, userService, msg)
		}

		if password == nil {
			msg = "Password cannot be empty"
			messages.PrintMessage(msg)
			editUserMenu(db, userService, msg)
		} else if len(password) < 6 {
			msg = "Password must be at least 6 characters"
			messages.PrintMessage(msg)
			editUserMenu(db, userService, msg)
		} else if strings.Contains(string(password), " ") {
			msg = "Password cannot contain spaces"
			messages.PrintMessage(msg)
			editUserMenu(db, userService, msg)
		}

		confirmPassword, err := terminal.HidePassword("Confirm new password (min 6 characters):")
		if err != nil {
			msg = "Error reading password"
			messages.PrintMessage(msg)
			editUserMenu(db, userService, msg)
		}

		if string(password) != string(confirmPassword) {
			msg = "Password and confirm password do not match"
			messages.PrintMessage(msg)
			editUserMenu(db, userService, msg)
		}

		user.Password = string(password)
	} else if updatePassword != "n" {
		msg = "Invalid input"
		messages.PrintMessage(msg)
		editUserMenu(db, userService, msg)
	}

	err = userService.UpdateUser(user)
	if err != nil {
		if strings.Contains(err.Error(), "users_email_key") {
			msg = "Email already exists"
		} else if strings.Contains(err.Error(), "users_username_key") {
			msg = "Username already exists"
		} else {
			msg = "Error updating user"
		}
		messages.PrintMessage(msg)
		editUserMenu(db, userService, msg)
	}

	ManageUserMenu(db, "Successfully updated user with ID "+strconv.Itoa(user.UserID))
}

func deleteUserMenu(db *sqlx.DB, userService *services.UserService, message string) {
	fmt.Println("=====================================")
	fmt.Println("Delete User")
	fmt.Println("=====================================")

	messages.PrintMessage(message)

	showUsers(userService)

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
