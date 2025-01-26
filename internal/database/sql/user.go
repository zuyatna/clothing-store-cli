package sql

import (
	"clothing-pair-project/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func (repository *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User
	query := `SELECT * FROM users`
	err := repository.db.Select(&users, query)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (repository *UserRepository) FindByUsername(username string) (models.User, error) {
	var user models.User
	query := `SELECT * FROM users WHERE username = $1`
	err := repository.db.Get(&user, query, username)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (repository *UserRepository) Add(user models.User) error {
	//TODO implement me
	panic("implement me")
}

func (repository *UserRepository) Update(user models.User) error {
	//TODO implement me
	panic("implement me")
}

func (repository *UserRepository) Delete(username string) error {
	//TODO implement me
	panic("implement me")
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}
