package handler

import (
	"clothing-pair-project/entity"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/jmoiron/sqlx"
)

type CategoryMethodHandler struct {
	db *sqlx.DB
}

func NewCategoryHandler(db *sqlx.DB) *CategoryMethodHandler {
	return &CategoryMethodHandler{db: db}
}

func (h *CategoryMethodHandler) Add(categories entity.Categories) error {
	query := `INSERT INTO categories (collection_id,name) VALUES ($1,$2)`
	_, err := h.db.Exec(query, categories.Collection_id, categories.Name)
	return err
}

func (h *CategoryMethodHandler) Delete(categories int) error {
	query := `DELETE FROM categories WHERE category_id = $1`
	_, err := h.db.Exec(query, categories)
	return err
}

func (h *CategoryMethodHandler) Update(categories entity.Categories) error {
	query := `UPDATE categories SET collection_id = $1, name = $2 WHERE category_id = $3`
	_, err := h.db.Exec(query, categories.Collection_id, categories.Name, categories.Category_id)
	return err
}

func (h *CategoryMethodHandler) Find(catID *int) ([]entity.Categories, error) {
	var categories []entity.Categories
	var query string
	var err error

	if catID == nil {
		query = `SELECT * FROM categories`
		err = h.db.Select(&categories, query)
	} else {
		query = `SELECT * FROM categories WHERE category_id = $1`
		err = h.db.Select(&categories, query, *catID)
	}

	if len(categories) == 0 {
		return nil, fmt.Errorf("no data found")
	}

	if err != nil {
		return nil, err
	}
	return categories, nil
}

func ShowDataCategory(namatable string, sizes []entity.Categories) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)

	fmt.Println(strings.Repeat("=", 40))
	fmt.Println(strings.Repeat(" ", 15) + namatable + strings.Repeat(" ", 15))
	fmt.Println(strings.Repeat("=", 40))
	_, _ = w.Write([]byte("ID\tCollectionID\tName\tCreatedAt\n"))
	_, _ = w.Write([]byte("--\t--\t--\t----\n"))

	for _, size := range sizes {
		_, _ = w.Write([]byte(
			fmt.Sprintf("%d\t%d\t%s\t%s\n", size.Category_id, size.Collection_id, size.Name, size.Created_at),
		))
	}

	_ = w.Flush()
	fmt.Println(strings.Repeat("=", 40))
}
