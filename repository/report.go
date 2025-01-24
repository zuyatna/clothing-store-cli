package repository

import "clothing-pair-project/entity"

type ReportRepository interface {
	ReportNeedRestock() ([]entity.ReportNeedRestock, error)
	ReportRevenue(year, month int) ([]entity.ReportRevenue, error)
}
