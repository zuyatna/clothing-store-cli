package handler

import (
	"clothing-pair-project/entity"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

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

func ShowDataNeedRestock(namatable string, datas []entity.ReportNeedRestock) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)

	fmt.Println(strings.Repeat("=", 40))
	fmt.Println(strings.Repeat(" ", 15) + namatable + strings.Repeat(" ", 15))
	fmt.Println(strings.Repeat("=", 40))
	_, _ = w.Write([]byte("Name\tCategory\tColor\tSize\tStock\n"))
	_, _ = w.Write([]byte("--\t--\t--\t--\t--\n"))

	for _, data := range datas {
		_, _ = w.Write([]byte(
			fmt.Sprintf("%s\t%s\t%s\t%s\t%d\n", data.Name, data.Category, data.Color, data.Size, data.Stock),
		))
	}

	_ = w.Flush()
	fmt.Println(strings.Repeat("=", 40))
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

func ShowDataNeedRevenue(namatable string, datas []entity.ReportRevenue) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)

	fmt.Println(strings.Repeat("=", 40))
	fmt.Println(strings.Repeat(" ", 15) + namatable + strings.Repeat(" ", 15))
	fmt.Println(strings.Repeat("=", 40))
	_, _ = w.Write([]byte("Year\tMonth\tRevenue\n"))
	_, _ = w.Write([]byte("--\t--\t--\n"))

	for _, data := range datas {
		_, _ = w.Write([]byte(
			fmt.Sprintf("%d\t%d\t%.2f\n", data.Year, data.Month, data.Revenue),
		))
	}

	_ = w.Flush()
	fmt.Println(strings.Repeat("=", 40))
}
