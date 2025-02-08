package menu

import (
	"clothing-pair-project/internal/database/sqlrepo"
	"clothing-pair-project/internal/models"
	"clothing-pair-project/internal/services"
	"clothing-pair-project/internal/utils/handler"
	"clothing-pair-project/internal/utils/key_input"
	"clothing-pair-project/internal/utils/messages"
	"clothing-pair-project/internal/utils/tables"
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

	userRepository := sqlrepo.NewUserQuery(db)
	userService := services.NewUserService(userRepository)

	var input string
	fmt.Print("Choose option: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		msg = "No input entered"
		ManageUserMenu(db, msg)
		return
	}

	switch input {
	case "1":
		findAllUsersMenu(db, userService, "", 5, 0)
		return
	case "2":
		findUserByUsername(db, userService, "")
		return
	case "3":
		addUserMenu(db, userService, "")
		return
	case "4":
		editUserMenu(db, userService, "")
		return
	case "5":
		deleteUserMenu(db, userService, "")
		return
	case "0":
		AdminMenu(db, "")
		return
	default:
		msg = "Invalid input"
		ManageUserMenu(db, msg)
		return
	}
}

func showUsers(userService *services.UserService, limit, offset int) (bool, bool) {
	table := tablewriter.NewWriter(os.Stdout)
	displayed := tables.UsersTablePresenter(table)
	userHandler := handler.NewUserHandler(displayed, userService)

	hasNext, hasPrev, err := userHandler.ShowAllUsers(limit, offset)
	if err != nil {
		fmt.Println("Error fetching users:", err)
		return false, false
	}

	return hasNext, hasPrev
}

func showUser(user models.User) {
	table := tablewriter.NewWriter(os.Stdout)
	displayed := tables.AddUserTablePresenter(table)
	displayed.DisplayAddUser(user)
}

func showUserByUsername(db *sqlx.DB, userService *services.UserService, msg string, username string) {
	table := tablewriter.NewWriter(os.Stdout)
	displayed := tables.UsersTablePresenter(table)
	userHandler := handler.NewUserHandler(displayed, userService)
	err := userHandler.ShowUserByUsername(username)
	if err != nil {
		msg = "Error fetching user by username"
		findUserByUsername(db, userService, msg)
	}
}

func findAllUsersMenu(db *sqlx.DB, userService *services.UserService, msg string, limit, offset int) {
	fmt.Println("=====================================")
	fmt.Println("Find All Users")
	fmt.Println("=====================================")

	hasNext, hasPrev := showUsers(userService, limit, offset)
	messages.PrintMessage(msg)

	if hasPrev {
		fmt.Println("Type A to Previous")
	}
	if hasNext {
		fmt.Println("Type D to Next")
	}
	fmt.Println("Type 0 to Back")
	fmt.Print("Choose option: ")

	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		msg = "No input entered"
		ManageUserMenu(db, msg)
		return
	}

	switch input {
	case "D", "d":
		if hasNext {
			offset += limit
		}
		findAllUsersMenu(db, userService, "", limit, offset)
		return
	case "A", "a":
		if hasPrev {
			offset -= limit
			if offset < 0 {
				offset = 0
			}
		}
		findAllUsersMenu(db, userService, "", limit, offset)
		return
	case "0":
		ManageUserMenu(db, "")
		return
	default:
		msg = "Invalid input"
		findAllUsersMenu(db, userService, msg, limit, offset)
		return
	}
}

func findUserByUsername(db *sqlx.DB, userService *services.UserService, msg string) {
	fmt.Println("=====================================")
	fmt.Println("Find User By Username")
	fmt.Println("=====================================")

	messages.PrintMessage(msg)

	username, err := key_input.Username()
	if err != nil {
		msg = err.Error()
		findUserByUsername(db, userService, msg)
		return
	}
	fmt.Println()

	showUserByUsername(db, userService, msg, username)

	key_input.BackMenu()
	ManageUserMenu(db, "")
}

func addUserMenu(db *sqlx.DB, userService *services.UserService, msg string) {
	fmt.Println("=====================================")
	fmt.Println("Add User")
	fmt.Println("=====================================")

	messages.PrintMessage(msg)

	username, err := key_input.Username()
	if err != nil {
		msg = err.Error()
		addUserMenu(db, userService, msg)
		return
	}

	email, err := key_input.Email()
	if err != nil {
		msg = err.Error()
		addUserMenu(db, userService, msg)
		return
	}

	password, err := key_input.Password()
	if err != nil {
		msg = err.Error()
		addUserMenu(db, userService, msg)
		return
	}

	role, err := key_input.Role(db)
	if err != nil {
		msg = err.Error()
		addUserMenu(db, userService, msg)
		return
	}

	user := models.User{
		Username: username,
		Email:    email,
		Password: password,
		Role:     role,
	}

	showUser(user)

	confirm, err := key_input.ConfirmAddUser()
	if err != nil {
		msg = err.Error()
		addUserMenu(db, userService, msg)
		return
	}

	if !confirm {
		msg = "User not added"
		addUserMenu(db, userService, msg)
		return
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
		return
	}

	msg = "User added successfully"
	ManageUserMenu(db, msg)
}

func editUserMenu(db *sqlx.DB, userService *services.UserService, msg string) {
	fmt.Println("=====================================")
	fmt.Println("Edit User")
	fmt.Println("=====================================")

	messages.PrintMessage(msg)

	userID, err := key_input.UserID()
	if err != nil {
		msg = err.Error()
		editUserMenu(db, userService, msg)
		return
	}

	userIDInt, err := strconv.Atoi(strings.TrimSpace(userID))
	if err != nil {
		msg = "User ID must be a number"
		editUserMenu(db, userService, msg)
		return
	}

	user, err := userService.GetUserByID(userIDInt)
	if err != nil {
		msg = "User not found"
		editUserMenu(db, userService, msg)
		return
	}

	fmt.Println("Current Username:", user.Username)
	editUsername, err := key_input.EditUsername(user.Username)
	if err != nil {
		msg = err.Error()
		editUserMenu(db, userService, msg)
		return
	}
	user.Username = editUsername

	fmt.Println("Current Email:", user.Email)
	editEmail, err := key_input.EditEmail(user.Email)
	if err != nil {
		msg = err.Error()
		editUserMenu(db, userService, msg)
		return
	}
	user.Email = editEmail

	editPassword, err := key_input.EditPassword(user.Password)
	if err != nil {
		msg = err.Error()
		editUserMenu(db, userService, msg)
		return
	}
	user.Password = editPassword

	editRole, err := key_input.EditRole(db, user.Role)
	if err != nil {
		msg = err.Error()
		editUserMenu(db, userService, msg)
		return
	}
	user.Role = editRole

	user = models.User{
		Username: editUsername,
		Email:    editEmail,
		Password: editPassword,
		Role:     editRole,
	}

	showUser(user)

	confirm, err := key_input.ConfirmEditUser()
	if err != nil {
		msg = err.Error()
		editUserMenu(db, userService, msg)
		return
	}

	if !confirm {
		msg = "User not updated"
		ManageUserMenu(db, msg)
		return
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
		editUserMenu(db, userService, msg)
		return
	}

	msg = "User updated successfully"
	ManageUserMenu(db, msg)
}

func deleteUserMenu(db *sqlx.DB, userService *services.UserService, msg string) {
	fmt.Println("=====================================")
	fmt.Println("Delete User")
	fmt.Println("=====================================")

	messages.PrintMessage(msg)

	userID, err := key_input.UserID()
	if err != nil {
		msg = err.Error()
		deleteUserMenu(db, userService, msg)
		return
	}

	userIDInt, err := strconv.Atoi(strings.TrimSpace(userID))
	if err != nil {
		msg = "Invalid User ID"
		deleteUserMenu(db, userService, msg)
		return
	}

	confirm, err := key_input.ConfirmDeleteUser()
	if err != nil {
		msg = err.Error()
		deleteUserMenu(db, userService, msg)
		return
	}

	if !confirm {
		msg = "User not deleted"
		ManageUserMenu(db, msg)
		return
	}

	err = userService.DeleteUser(userIDInt)
	if err != nil {
		msg = "Error deleting user"
		deleteUserMenu(db, userService, msg)
		return
	}

	msg = "User deleted successfully"
	ManageUserMenu(db, msg)
}
