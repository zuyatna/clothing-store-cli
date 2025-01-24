package entity

type Color struct {
	ColorID    int    `db:"color_id"`
	Name       string `db:"name"`
	Created_at string `db:"created_at"`
}
