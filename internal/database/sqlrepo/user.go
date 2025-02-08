package sqlrepo

import (
	"clothing-pair-project/internal/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UserQuery struct {
	db *sqlx.DB
}

func NewUserQuery(db *sqlx.DB) *UserQuery {
	return &UserQuery{db: db}
}

func (repository *UserQuery) FindAll(limit, offset int) ([]models.User, error) {
	var users []models.User
	query := "SELECT * FROM users ORDER BY user_id ASC LIMIT $1 OFFSET $2"
	err := repository.db.Select(&users, query, limit, offset)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (repository *UserQuery) FindByID(id int) (models.User, error) {
	var user models.User
	query := "SELECT * FROM users WHERE user_id = $1"
	err := repository.db.Get(&user, query, id)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (repository *UserQuery) FindByUsername(username string) (models.User, error) {
	var user models.User
	query := "SELECT * FROM users WHERE username = $1"
	err := repository.db.Get(&user, query, username)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (repository *UserQuery) Add(user models.User) error {
	nextID, err := repository.GetNextID()
	if err != nil {
		return fmt.Errorf("error getting next ID: %v", err)
	}

	query := `INSERT INTO users (user_id, username, email, password, role, created_at, active) 
              VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, true)`

	_, err = repository.db.Exec(query, nextID, user.Username, user.Email, user.Password, user.Role)
	if err != nil {
		return err
	}

	return nil
}

func (repository *UserQuery) Update(user models.User) error {
	query := "UPDATE users SET username = $1, email = $2, password = $3, role = $4 WHERE user_id = $5"
	_, err := repository.db.Exec(query, user.Username, user.Email, user.Password, user.Role, user.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (repository *UserQuery) Delete(id int) error {
	query := "UPDATE users SET active = false WHERE user_id = $1"
	_, err := repository.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (repository *UserQuery) GetNextID() (int, error) {
	var id int

	createSeq := `DO $$ 
    BEGIN
        CREATE SEQUENCE IF NOT EXISTS "Users_UserID_seq";
    END $$;`

	_, err := repository.db.Exec(createSeq)
	if err != nil {
		return 0, err
	}

	query := `SELECT COALESCE(
        (SELECT MAX(user_id) + 1 FROM users), 
        nextval('"Users_UserID_seq"')
    )`

	err = repository.db.Get(&id, query)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (repository *UserQuery) EnumRole() (role string, err error) {
	query := "SELECT enum_range(NULL::role)"
	err = repository.db.Get(&role, query)
	if err != nil {
		return "role not found", err
	}

	return role, nil
}
