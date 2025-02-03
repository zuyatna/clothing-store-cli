package sqlrepo

import "github.com/jmoiron/sqlx"

type ColorRepository struct {
	db *sqlx.DB
}

func NewColorRepository(db *sqlx.DB) *ColorRepository {
	return &ColorRepository{db: db}
}

func (repository *ColorRepository) FindAll() ([]string, error) {
	var colors []string
	query := "SELECT color FROM colors"
	err := repository.db.Select(&colors, query)
	if err != nil {
		return nil, err
	}

	return colors, nil
}

func (repository *ColorRepository) FindByID(id int) (string, error) {
	var color string
	query := "SELECT color FROM colors WHERE color_id = $1"
	err := repository.db.Get(&color, query, id)
	if err != nil {
		return "", err
	}

	return color, nil
}

func (repository *ColorRepository) Add(color string) error {
	query := `INSERT INTO colors (color) VALUES ($1)`
	_, err := repository.db.Exec(query, color)
	if err != nil {
		return err
	}

	return nil
}

func (repository *ColorRepository) Update(color string) error {
	query := "UPDATE colors SET color = $1 WHERE color_id = $2"
	_, err := repository.db.Exec(query, color)
	if err != nil {
		return err
	}

	return nil
}

func (repository *ColorRepository) Delete(id int) error {
	query := "DELETE FROM colors WHERE color_id = $1"
	_, err := repository.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
