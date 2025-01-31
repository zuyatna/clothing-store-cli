package models

import "time"

type Category struct {
	CategoryID int       `json:"category_id" db:"category_id"`
	Name       string    `json:"name" db:"name"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}
