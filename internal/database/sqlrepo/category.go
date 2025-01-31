package sqlrepo

import (
	"clothing-pair-project/internal/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type CategoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (repository *CategoryRepository) FindAll() ([]models.Category, error) {
	var categories []models.Category
	query := "SELECT * FROM categories ORDER BY category_id ASC"
	err := repository.db.Select(&categories, query)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (repository *CategoryRepository) FindByID(id int) (models.Category, error) {
	var category models.Category
	query := "SELECT * FROM categories WHERE category_id = $1"
	err := repository.db.Get(&category, query, id)
	if err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func (repository *CategoryRepository) Add(category models.Category) error {
	nextID, err := repository.GetNextID()
	if err != nil {
		return fmt.Errorf("error getting next ID: %v", err)
	}

	query := `INSERT INTO categories (category_id, name)
    		  VALUES ($1, $2)`

	_, err = repository.db.Exec(query, nextID, category.Name)
	if err != nil {
		return err
	}

	return nil
}

func (repository *CategoryRepository) Update(category models.Category) error {
	query := "UPDATE categories SET name = $1 WHERE category_id = $2"
	_, err := repository.db.Exec(query, category.Name, category.CategoryID)
	if err != nil {
		return err
	}

	return nil
}

func (repository *CategoryRepository) Delete(id int) error {
	query := "DELETE FROM categories WHERE category_id = $1"
	_, err := repository.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (repository *CategoryRepository) GetNextID() (int, error) {
	var id int

	createSeq := `DO $$
	BEGIN
		CREATE SEQUENCE IF NOT EXISTS "Category_CategoryID_seq";
	END $$;`

	_, err := repository.db.Exec(createSeq)
	if err != nil {
		return 0, err
	}

	query := `SELECT COALESCE(MAX(category_id), 0) + 1 FROM categories`

	err = repository.db.Get(&id, query)
	if err != nil {
		return 0, err
	}

	return id, nil
}
