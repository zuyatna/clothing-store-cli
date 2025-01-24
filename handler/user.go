package handler

import (
	"clothing-pair-project/entity"

	"github.com/jmoiron/sqlx"
)

type UserHandler struct {
	db *sqlx.DB
}

func NewUserHandler(db *sqlx.DB) *UserHandler {
	return &UserHandler{db: db}
}

func (h *UserHandler) FindAll() ([]entity.User, error) {
	var users []entity.User
	query := `SELECT * FROM users`
	err := h.db.Select(&users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (h *UserHandler) FindByID(userID int) (entity.User, error) {
	var user entity.User
	query := `SELECT * FROM users WHERE user_id = $1`
	err := h.db.Get(&user, query, userID)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (h *UserHandler) FindByUsername(username string) (entity.User, error) {
	var user entity.User
	query := `SELECT user_id, username, email, password, role FROM users WHERE username = $1`
	err := h.db.Get(&user, query, username)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (h *UserHandler) Add(user entity.User) error {
	nextID, err := h.GetNextID()
	if err != nil {
		return err
	}

	query := `INSERT INTO users (user_id, username, email, password, role, created_at) 
              VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP)`
	_, err = h.db.Exec(query, nextID, user.Username, user.Email, user.Password, user.Role)
	return err
}

func (h *UserHandler) Delete(userID int) error {
	query := `DELETE FROM users WHERE user_id = $1`
	_, err := h.db.Exec(query, userID)
	return err
}

func (h *UserHandler) Update(user entity.User) error {
	query := `UPDATE users SET username = $1, email = $2, password = $3, role = $4 
              WHERE user_id = $5`
	_, err := h.db.Exec(query, user.Username, user.Email, user.Password, user.Role, user.UserID)
	return err
}

func (h *UserHandler) ResetIncrement() error {
	query := `ALTER SEQUENCE "User_UserID_seq" RESTART WITH 1`
	_, err := h.db.Exec(query)
	return err
}

func (h *UserHandler) GetNextID() (int, error) {
	var nextID int
	query := `SELECT setval('"Users_UserID_seq"', (SELECT MAX(user_id)+1 FROM users));`
	err := h.db.Get(&nextID, query)
	return nextID, err
}
