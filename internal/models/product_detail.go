package models

import "time"

type ProductDetail struct {
	ProductDetailID int       `json:"product_detail_id" db:"product_detail_id"`
	ProductID       int       `json:"product_id" db:"product_id"`
	ColorID         int       `json:"color_id" db:"color_id"`
	SizeID          int       `json:"size_id" db:"size_id"`
	Stock           int       `json:"stock" db:"stock"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
}
