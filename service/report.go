package service

import (
	"clothing-pair-project/entity"
	"clothing-pair-project/repository"
)

type ReportService struct {
	reportRepository repository.ReportRepository
}

func NewReportService(reportRepository repository.ReportRepository) *ReportService {
	return &ReportService{reportRepository}
}

func (service *ReportService) ReportNeedRestock() ([]entity.ReportNeedRestock, error) {
	report, err := service.reportRepository.ReportNeedRestock()
	if err != nil {
		return report, err
	}
	return report, nil
}

func (service *ReportService) ReportRevenue(year, month int) ([]entity.ReportRevenue, error) {
	report, err := service.reportRepository.ReportRevenue(year, month)
	if err != nil {
		return report, err
	}
	return report, nil
}
