package repository

import "clothing-pair-project/internal/models"

type UserRepository interface {
	FindAll() ([]models.User, error)
	FindByUsername(username string) (models.User, error)
	Add(user models.User) error
	Update(user models.User) error
	Delete(username string) error
}
