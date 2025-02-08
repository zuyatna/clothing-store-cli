package handler

import (
	"clothing-pair-project/internal/models"
	"clothing-pair-project/internal/utils/interfaces"
	"fmt"
)

type UserHandler struct {
	userDisplay interfaces.UserDisplay
	userFetcher interfaces.UserFetcher
}

func NewUserHandler(userDisplay interfaces.UserDisplay, userFetcher interfaces.UserFetcher) *UserHandler {
	return &UserHandler{
		userDisplay: userDisplay,
		userFetcher: userFetcher,
	}
}

func (h *UserHandler) ShowAllUsers(limit, offset int) (bool, bool, error) {
	users, err := h.userFetcher.GetAllUsers(limit+1, offset)
	if err != nil {
		return false, false, fmt.Errorf("error fetching all users: %w", err)
	}

	if len(users) == 0 {
		return false, false, fmt.Errorf("no users found")
	}

	displayUsers := users
	if len(users) > limit {
		displayUsers = users[:limit]
	}

	h.userDisplay.DisplayUsers(displayUsers)

	hasNext := len(users) > limit
	hasPrev := offset > 0

	return hasNext, hasPrev, nil
}

func (h *UserHandler) ShowUserByUsername(username string) error {
	user, err := h.userFetcher.GetUserByUsername(username)
	if err != nil {
		return fmt.Errorf("error fetching user by username: %w", err)
	}

	h.userDisplay.DisplayUsers([]models.User{user})
	return nil
}
