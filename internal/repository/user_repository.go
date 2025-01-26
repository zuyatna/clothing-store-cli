package repository

import "clothing-pair-project/internal/models"

type UserRepository interface {
	FindAll() ([]models.User, error)
	FindByID(id int) (models.User, error)
	FindByUsername(username string) (models.User, error)
	Add(user models.User) error
	Update(user models.User) error
	Delete(id int) error
	EnumRole() (string, error)
}
