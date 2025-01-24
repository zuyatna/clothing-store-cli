package entity

import "time"

type Products struct {
	Product_Id  int
	Category_Id int
	Color_Id    int
	Size_Id     int
	Name        string
	Price       float32
	Stock       int
	Description string
	Image       string
	Created_At  string
}

type ShowDataProducts struct {
	Product_Id  int       `db:"product_id"`
	Category_Id string    `db:"category"`
	Color_Id    string    `db:"color"`
	Size_Id     string    `db:"size"`
	Name        string    `db:"name"`
	Price       float32   `db:"price"`
	Stock       int       `db:"stock"`
	Description string    `db:"description"`
	Image       string    `db:"image"`
	Created_At  time.Time `db:"created_at"`
}
