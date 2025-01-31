package models

import "time"

type Color struct {
	ColorID   int       `json:"color_id" json:"color_id"`
	Name      string    `json:"name" json:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
