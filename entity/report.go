package entity

type ReportNeedRestock struct {
	Name     string `db:"name"`
	Category string `db:"category"`
	Color    string `db:"color"`
	Size     string `db:"size"`
	Stock    int    `db:"stock"`
}

type ReportRevenue struct {
	Year    int     `db:"year"`
	Month   int     `db:"month"`
	Revenue float32 `db:"revenue"`
}

type TotalPurchase struct {
	ProductID   int     `db:"product_id"`
	ProductName string  `db:"product_name"`
	Count       int     `db:"count"`
	Total       float64 `db:"total"`
}

type TotalUser struct {
	Total int `db:"total"`
}
