package service

import (
	"clothing-pair-project/entity"
	"clothing-pair-project/repository"
	"log"
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

func (service *ReportService) TotalPurchase() ([]entity.TotalPurchase, error) {
	report, err := service.reportRepository.TotalPurchase()
	if err != nil {
		log.Printf("TotalPurchase error: %v", err)
		return nil, err
	}
	return report, nil
}

func (service *ReportService) TotalUser() ([]entity.TotalUser, error) {
	report, err := service.reportRepository.TotalUser()
	if err != nil {
		return report, err
	}
	return report, nil
}
