package entity

import "time"

type Collection struct {
	CollectionID int       `json:"collection_id" db:"collection_id"`
	Name         string    `json:"name" db:"name"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}
