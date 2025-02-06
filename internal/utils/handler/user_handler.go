package handler

import (
	"clothing-pair-project/internal/models"
	"clothing-pair-project/internal/utils/interfaces"
	"fmt"
)

type UserHandler struct {
	userService interfaces.UserFetcher
	userDisplay interfaces.UserDisplay
}

func NewUserHandler(service interfaces.UserFetcher, display interfaces.UserDisplay) *UserHandler {
	return &UserHandler{
		userService: service,
		userDisplay: display,
	}
}

func (h *UserHandler) ShowAllUsers() error {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		return fmt.Errorf("error fetching all users: %w", err)
	}

	h.userDisplay.DisplayUsers(users)
	return nil
}

func (h *UserHandler) ShowUserByUsername(username string) error {
	user, err := h.userService.GetUserByUsername(username)
	if err != nil {
		return fmt.Errorf("error fetching user by username: %w", err)
	}

	h.userDisplay.DisplayUsers([]models.User{user})
	return nil
}
