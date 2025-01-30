package models

import (
	"database/sql"
	"time"
)

type Product struct {
	ProductID   int            `json:"product_id" db:"product_id"`
	CategoryID  int            `json:"category_id" db:"category_id"`
	Name        string         `json:"name" db:"name"`
	Price       float64        `json:"price" db:"price"`
	Description sql.NullString `json:"description" db:"description"`
	Images      sql.NullString `json:"images" db:"images"`
	Type        string         `json:"type" db:"type"`
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
}
