package interfaces

import (
	"clothing-pair-project/internal/models"
)

type UserDisplay interface {
	DisplayAllUser(users []models.User)
}

type UserFetcher interface {
	GetAllUsers() ([]models.User, error)
}
