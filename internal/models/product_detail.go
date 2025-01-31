package models

import "time"

type ProductDetail struct {
	ProductDetailID int       `json:"product_detail_id" db:"product_detail_id"`
	ProductID       int       `json:"product_id" db:"product_id"`
	ColorID         int       `json:"color_id" db:"color_id"`
	Color           string    `json:"color" db:"color"`
	SizeID          int       `json:"size_id" db:"size_id"`
	Size            string    `json:"size" db:"size"`
	Stock           int       `json:"stock" db:"stock"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
}

type ProductDetailRequest struct {
	ProductID int    `json:"product_id" db:"product_id"`
	Color     string `json:"color" db:"color"`
	Size      string `json:"size" db:"size"`
	Stock     int    `json:"stock" db:"stock"`
}
