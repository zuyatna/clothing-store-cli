package handler

import (
	"clothing-pair-project/entity"

	"github.com/jmoiron/sqlx"
)

type ColorHandler struct {
	db *sqlx.DB
}

func NewColorHandler(db *sqlx.DB) *ColorHandler {
	return &ColorHandler{db: db}
}

func (h *ColorHandler) FindAll() ([]entity.Color, error) {
	var colors []entity.Color
	query := `SELECT color_id, name FROM colors`
	err := h.db.Select(&colors, query)
	if err != nil {
		return nil, err
	}
	return colors, nil
}

func (h *ColorHandler) FindByID(colorID int) (entity.Color, error) {
	var color entity.Color
	query := `SELECT color_id, name FROM colors WHERE color_id = $1`
	err := h.db.Get(&color, query, colorID)
	if err != nil {
		return entity.Color{}, err
	}
	return color, nil
}

func (h *ColorHandler) Add(color entity.Color) error {
	query := `INSERT INTO colors (name) VALUES ($1)`
	_, err := h.db.Exec(query, color.Name)
	return err
}

func (h *ColorHandler) Delete(colorID int) error {
	query := `DELETE FROM colors WHERE color_id = $1`
	_, err := h.db.Exec(query, colorID)
	return err
}

func (h *ColorHandler) Update(color entity.Color) error {
	query := `UPDATE colors SET name = $1 WHERE color_id = $2`
	_, err := h.db.Exec(query, color.Name, color.ColorID)
	return err
}

func (h *ColorHandler) ResetIncrement() error {
	query := `ALTER SEQUENCE "Color_ColorID_seq" RESTART WITH 1`
	_, err := h.db.Exec(query)
	return err
}
