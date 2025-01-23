package handler

import (
	"clothing-pair-project/entity"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type SizesMethodHandler struct {
	db *sqlx.DB
}

func NewSizesHandler(db *sqlx.DB) *SizesMethodHandler {
	return &SizesMethodHandler{db: db}
}

func (h *SizesMethodHandler) Add(sizes entity.Sizes) error {
	query := `INSERT INTO sizes (name) VALUES ($1)`
	_, err := h.db.Exec(query, sizes.Name)
	return err
}

func (h *SizesMethodHandler) Delete(sizes int) error {
	query := `DELETE FROM sizes WHERE size_id = $1`
	_, err := h.db.Exec(query, sizes)
	return err
}

func (h *SizesMethodHandler) Update(sizes entity.Sizes) error {
	query := `UPDATE sizes SET name = $1 WHERE size_id = $2`
	_, err := h.db.Exec(query, sizes.Name, sizes.Size_id)
	return err
}

func (h *SizesMethodHandler) Find(sizeID *int) ([]entity.Sizes, error) {
	var sizes []entity.Sizes
	var query string
	var err error

	if sizeID == nil {
		query = `SELECT * FROM sizes`
		err = h.db.Select(&sizes, query)
	} else {
		query = `SELECT * FROM sizes WHERE size_id = $1`
		err = h.db.Select(&sizes, query, *sizeID)
	}

	if len(sizes) == 0 {
		return nil, fmt.Errorf("no data found")
	}

	if err != nil {
		return nil, err
	}
	return sizes, nil
}
