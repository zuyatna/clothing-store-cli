package interfaces

import (
	"clothing-pair-project/internal/models"
)

type UserDisplay interface {
	DisplayUsers(users []models.User)
}

type UserFetcher interface {
	GetAllUsers(limit, offset int) ([]models.User, error)
	GetUserByUsername(username string) (models.User, error)
}
