package models

import "time"

type Product struct {
	ProductID   int       `json:"product_id" db:"product_id"`
	CategoryID  int       `json:"category_id" db:"category_id"`
	Name        string    `json:"name" db:"name"`
	Price       float64   `json:"price" db:"price"`
	Description string    `json:"description" db:"description"`
	Image       string    `json:"image" db:"image"`
	Type        string    `json:"type" db:"type"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
