package repository

import "clothing-pair-project/entity"

type UserRepository interface {
	FindAll() ([]entity.User, error)
	FindByID(userID int) (entity.User, error)
	FindByUsername(username string) (entity.User, error)
	Add(user entity.User) error
	Update(user entity.User) error
	Delete(userID int) error
	ResetIncrement() error
}
