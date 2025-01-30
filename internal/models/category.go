package models

type Category struct {
	CategoryID int    `json:"category_id" db:"category_id"`
	Name       string `json:"name" db:"name"`
	CreatedAt  string `json:"created_at" db:"created_at"`
}
