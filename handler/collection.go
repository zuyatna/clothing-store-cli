package handler

import (
	"clothing-pair-project/entity"

	"github.com/jmoiron/sqlx"
)

type CollectionHandler struct {
	db *sqlx.DB
}

func NewCollectionHandler(db *sqlx.DB) *CollectionHandler {
	return &CollectionHandler{db: db}
}

func (h *CollectionHandler) FindByID(collectionID int) (entity.Collection, error) {
	var collection entity.Collection
	query := `SELECT * FROM collections WHERE collection_id = $1`
	err := h.db.Get(&collection, query, collectionID)
	if err != nil {
		return entity.Collection{}, err
	}
	return collection, nil
}

func (h *CollectionHandler) FindAll() ([]entity.Collection, error) {
	var collections []entity.Collection
	query := `SELECT collection_id, name, created_at FROM collections`
	err := h.db.Select(&collections, query)
	if err != nil {
		return nil, err
	}
	return collections, nil
}

func (h *CollectionHandler) Add(collection entity.Collection) error {
	query := `INSERT INTO collections (collection_id, name, created_at) 
              VALUES (COALESCE((SELECT MAX(collection_id) + 1 FROM collections), 1), $1, CURRENT_TIMESTAMP)`
	_, err := h.db.Exec(query, collection.Name)
	return err
}

func (h *CollectionHandler) Update(collection entity.Collection) error {
	query := `UPDATE collections SET name = $1, created_at = CURRENT_TIMESTAMP WHERE collection_id = $2`
	_, err := h.db.Exec(query, collection.Name, collection.CollectionID)
	return err
}

func (h *CollectionHandler) Delete(collectionID int) error {
	query := `DELETE FROM collections WHERE collection_id = $1`
	_, err := h.db.Exec(query, collectionID)
	return err
}

func (h *CollectionHandler) ResetIncrement() error {
	query := `ALTER SEQUENCE "Collection_CollectionID_seq" RESTART WITH 1`
	_, err := h.db.Exec(query)
	return err
}
