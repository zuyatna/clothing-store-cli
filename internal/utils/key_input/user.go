package key_input

import (
	"bufio"
	"clothing-pair-project/internal/database/sqlrepo"
	"clothing-pair-project/internal/helper"
	"clothing-pair-project/internal/utils/terminal"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/jmoiron/sqlx"
)

func UserID() (string, error) {
	fmt.Print("Enter User ID: ")
	reader := bufio.NewReader(os.Stdin)
	userID, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("error reading input: %w", err)
	}

	if userID == "" {
		return "", fmt.Errorf("user ID cannot be empty")
	} else if strings.Contains(userID, " ") {
		return "", fmt.Errorf("user ID cannot contain spaces")
	}

	return userID, nil
}

func Username() (string, error) {
	fmt.Print("Enter username: ")
	reader := bufio.NewReader(os.Stdin)
	username, err := reader.ReadString('\n')
	username = strings.TrimSpace(username)
	if err != nil {
		return "", fmt.Errorf("error reading input: %w", err)
	}

	if username == "" {
		fmt.Println("Username cannot be empty")
		return "", fmt.Errorf("username cannot be empty")
	} else if strings.Contains(username, " ") {
		fmt.Println("Username cannot contain spaces")
		return "", fmt.Errorf("username cannot contain spaces")
	}

	return username, nil
}

func Email() (string, error) {
	fmt.Print("Enter email: ")
	reader := bufio.NewReader(os.Stdin)
	email, err := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	if err != nil {
		return "", fmt.Errorf("error reading input: %w", err)
	}

	if email == "" {
		return "", fmt.Errorf("email cannot be empty")
	} else if strings.Contains(email, " ") {
		return "", fmt.Errorf("email cannot contain spaces")
	} else if !strings.Contains(email, "@") {
		return "", fmt.Errorf("email must contain @")
	}

	return email, nil
}

func Password() (string, error) {
	password, err := terminal.HidePassword("Enter Password: ")
	if err != nil {
		return "", fmt.Errorf("error reading input: %w", err)
	}

	if password == nil {
		return "", fmt.Errorf("password cannot be empty")
	} else if strings.Contains(string(password), " ") {
		return "", fmt.Errorf("password cannot contain spaces")
	}

	confirmPassword, err := terminal.HidePassword("Confirm Password:")
	if err != nil {
		return "", fmt.Errorf("error reading input: %w", err)
	}

	if string(password) != string(confirmPassword) {
		return "", fmt.Errorf("passwords do not match")
	}

	hashedPassword, err := helper.HashPassword(string(password))
	if err != nil {
		return "", fmt.Errorf("error hashing password: %w", err)
	}

	return hashedPassword, nil
}

func Role(db *sqlx.DB) (string, error) {
	userRepository := sqlrepo.NewUserQuery(db)
	enumRange, err := userRepository.EnumRole()
	if err != nil {
		return "", fmt.Errorf("error fetching enum role: %w", err)
	}

	roles := strings.Fields(strings.ReplaceAll(strings.Trim(enumRange, "{}"), ",", " "))

	fmt.Printf("Role %s: ", enumRange)
	reader := bufio.NewReader(os.Stdin)
	role, err := reader.ReadString('\n')
	role = strings.TrimSpace(role)
	if err != nil {
		return "", fmt.Errorf("error reading input: %w", err)
	}

	if role == "" {
		return "", fmt.Errorf("role cannot be empty")
	}

	if !slices.Contains(roles, role) {
		return "", fmt.Errorf("role must be one of %s", enumRange)
	}

	return role, nil
}

func ConfirmAddUser() (bool, error) {
	fmt.Print("Add user? (y/n): ")
	reader := bufio.NewReader(os.Stdin)
	confirm, err := reader.ReadString('\n')
	confirm = strings.TrimSpace(confirm)
	if err != nil {
		return false, fmt.Errorf("error reading input: %w", err)
	}

	if confirm == "y" {
		return true, nil
	}

	return false, nil
}

func EditUsername(currentUsername string) (string, error) {
	fmt.Print("Do you want to update username? (y/n): ")
	reader := bufio.NewReader(os.Stdin)
	updateUsername, err := reader.ReadString('\n')
	updateUsername = strings.TrimSpace(updateUsername)
	if err != nil {
		return currentUsername, fmt.Errorf("error reading input: %w", err)
	}

	if updateUsername == "y" {
		username, err := Username()
		if err != nil {
			return currentUsername, fmt.Errorf("error reading input: %w", err)
		}

		return username, nil
	} else if updateUsername != "n" {
		return "", fmt.Errorf("invalid input")
	}

	return currentUsername, nil
}

func EditEmail(currentEmail string) (string, error) {
	fmt.Print("Do you want to update email? (y/n): ")
	reader := bufio.NewReader(os.Stdin)
	updateEmail, err := reader.ReadString('\n')
	updateEmail = strings.TrimSpace(updateEmail)
	if err != nil {
		return currentEmail, fmt.Errorf("error reading input: %w", err)
	}

	if updateEmail == "y" {
		email, err := Email()
		if err != nil {
			return currentEmail, fmt.Errorf("error reading input: %w", err)
		}

		return email, nil
	} else if updateEmail != "n" {
		return currentEmail, fmt.Errorf("invalid input")
	}

	return currentEmail, nil
}

func EditPassword(currentPassword string) (string, error) {
	fmt.Print("Do you want to update password? (y/n): ")
	reader := bufio.NewReader(os.Stdin)
	updatePassword, err := reader.ReadString('\n')
	updatePassword = strings.TrimSpace(updatePassword)
	if err != nil {
		return currentPassword, fmt.Errorf("error reading input: %w", err)
	}

	if updatePassword == "y" {
		password, err := Password()
		if err != nil {
			return currentPassword, fmt.Errorf("error reading input: %w", err)
		}

		return password, nil
	} else if updatePassword != "n" {
		return currentPassword, fmt.Errorf("invalid input")
	}

	return currentPassword, nil
}

func EditRole(db *sqlx.DB, currentRole string) (string, error) {
	fmt.Print("Do you want to update role? (y/n): ")
	reader := bufio.NewReader(os.Stdin)
	updateRole, err := reader.ReadString('\n')
	updateRole = strings.TrimSpace(updateRole)
	if err != nil {
		return currentRole, fmt.Errorf("error reading input: %w", err)
	}

	if updateRole == "y" {
		role, err := Role(db)
		if err != nil {
			return currentRole, fmt.Errorf("error reading input: %w", err)
		}

		return role, nil
	} else if updateRole != "n" {
		return currentRole, fmt.Errorf("invalid input")
	}

	return currentRole, nil
}
