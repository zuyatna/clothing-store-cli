package sqlrepo

import "github.com/jmoiron/sqlx"

type SizeRepository struct {
	db *sqlx.DB
}

func NewSizeRepository(db *sqlx.DB) *SizeRepository {
	return &SizeRepository{db: db}
}

func (repository *SizeRepository) FindAll() ([]string, error) {
	var sizes []string
	query := "SELECT size FROM sizes"
	err := repository.db.Select(&sizes, query)
	if err != nil {
		return nil, err
	}

	return sizes, nil
}

func (repository *SizeRepository) FindByID(id int) (string, error) {
	var size string
	query := "SELECT size FROM sizes WHERE size_id = $1"
	err := repository.db.Get(&size, query, id)
	if err != nil {
		return "", err
	}

	return size, nil
}

func (repository *SizeRepository) Add(size string) error {
	query := `INSERT INTO sizes (size) VALUES ($1)`
	_, err := repository.db.Exec(query, size)
	if err != nil {
		return err
	}

	return nil
}

func (repository *SizeRepository) Update(size string) error {
	query := "UPDATE sizes SET size = $1 WHERE size_id = $2"
	_, err := repository.db.Exec(query, size)
	if err != nil {
		return err
	}

	return nil
}

func (repository *SizeRepository) Delete(id int) error {
	query := "DELETE FROM sizes WHERE size_id = $1"
	_, err := repository.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
