package models

import "time"

type Size struct {
	SizeID    int       `json:"size_id" db:"size_id"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
