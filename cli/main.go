package main

import (
	"clothing-pair-project/config"
	"clothing-pair-project/handler"
	"clothing-pair-project/service"
	"log"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect db", err.Error())
	}
	defer db.Close()

	reportMethodHandler := handler.NewReportHandler(db)
	reportMethodService := service.NewReportService(reportMethodHandler)

	// sizeId := 1
	needRestock, err := reportMethodService.ReportRevenue(2025, 1)
	if err != nil {
		log.Fatal("Failed to find data : ", err.Error())
	}
	handler.ShowDataNeedRevenue("Revenue", needRestock)
}
