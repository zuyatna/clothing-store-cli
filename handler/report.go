package handler

import (
	"clothing-pair-project/entity"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ReportHandler struct {
	db *sqlx.DB
}

func NewReportHandler(db *sqlx.DB) *ReportHandler {
	return &ReportHandler{db: db}
}

func (h *ReportHandler) ReportNeedRestock() ([]entity.ReportNeedRestock, error) {
	var report []entity.ReportNeedRestock
	query := `SELECT 
				P.NAME,
				C.NAME AS category,
				CO.name as color,
				S.name as size,
				P.stock
			FROM
				products P,
				categories C,
				colors CO,
				sizes S
			WHERE
				P.category_id = C.category_id 
				AND P.color_id = CO.color_id 
				AND P.size_id = S.size_id
				AND P.stock < 10 
			ORDER BY
				P.stock ASC`
	err := h.db.Select(&report, query)
	if err != nil {
		return nil, err
	}
	return report, nil
}

func (h *ReportHandler) ReportRevenue(year, month int) ([]entity.ReportRevenue, error) {
	var report []entity.ReportRevenue
	query := `SELECT 
				EXTRACT(YEAR FROM date) AS year,
				EXTRACT(MONTH FROM date) AS month,
				SUM (total) AS revenue 
			FROM
				purchases 
			WHERE
				EXTRACT (YEAR FROM date) = $1 
				AND EXTRACT (MONTH FROM date) = $2
			GROUP BY 
				EXTRACT(YEAR FROM date), 
				EXTRACT(MONTH FROM date)`
	err := h.db.Select(&report, query, year, month)
	if err != nil {
		return nil, err
	}
	return report, nil
}

func (h *ReportHandler) TotalPurchase() ([]entity.TotalPurchase, error) {
	var report []entity.TotalPurchase
	query := `SELECT 
				pd.product_id,
				p.name as product_name,
				COUNT(pd.product_id) as count,
				COALESCE(SUM(pd.sub_total), 0) as total
			FROM 
				purchase_details pd
				JOIN products p ON pd.product_id = p.product_id
			GROUP BY 
				pd.product_id, p.name
			ORDER BY
				pd.product_id ASC`

	if err := h.db.Select(&report, query); err != nil {
		return nil, fmt.Errorf("error getting total purchase: %v", err)
	}

	return report, nil
}

func (h *ReportHandler) TotalUser() ([]entity.TotalUser, error) {
	var report []entity.TotalUser
	query := `SELECT 
				COUNT(user_id) as total
			FROM 
				users`

	err := h.db.Select(&report, query)
	if err != nil {
		return nil, err
	}
	return report, nil
}
