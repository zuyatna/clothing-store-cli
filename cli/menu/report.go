package menu

import (
	"clothing-pair-project/handler"
	"clothing-pair-project/service"
	"fmt"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/olekukonko/tablewriter"
)

func ReportMenu(db *sqlx.DB) {
	for {
		fmt.Println()
		fmt.Println("=====================================")
		fmt.Println("Report Menu")
		fmt.Println("1. Report Need Restock")
		fmt.Println("2. Report Revenue")
		fmt.Println("3. Total Purchase")
		fmt.Println("4. Total User")
		fmt.Println("0. Back")
		fmt.Println("=====================================")

		var input int
		fmt.Print("Choose option: ")
		fmt.Scanln(&input)

		reportHandler := handler.NewReportHandler(db)
		reportService := service.NewReportService(reportHandler)

		switch input {
		case 1:
			reportNeedRestockMenu(reportService)
		case 2:
			reportRevenueMenu(reportService)
		case 3:
			totalPurchaseMenu(reportService)
		case 4:
			totalUserMenu(reportService)
		case 0:
			return
		default:
			fmt.Println("Invalid input")
		}
	}
}

func reportNeedRestockMenu(reportService *service.ReportService) {
	fmt.Println()
	fmt.Println("=====================================")
	fmt.Println("Report Need Restock")
	fmt.Println("=====================================")

	report, err := reportService.ReportNeedRestock()
	if err != nil {
		fmt.Println("Failed to get report need restock")
		return
	}

	if len(report) == 0 {
		fmt.Println("No data")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Product Name", "Category", "Color", "Size", "Stock"})
	for _, r := range report {
		table.Append([]string{r.Name, r.Category, r.Color, r.Size, strconv.Itoa(r.Stock)})
	}
	table.Render()

	fmt.Println()
}

func reportRevenueMenu(reportService *service.ReportService) {
	fmt.Println()
	fmt.Println("=====================================")
	fmt.Println("Report Revenue")
	fmt.Println("=====================================")

	var year, month int
	fmt.Print("Year: ")
	fmt.Scanln(&year)
	fmt.Print("Month: ")
	fmt.Scanln(&month)

	report, err := reportService.ReportRevenue(year, month)
	if err != nil {
		fmt.Println("Failed to get report revenue")
		return
	}

	if len(report) == 0 {
		fmt.Println("No data")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Year", "Month", "Revenue"})
	for _, r := range report {
		table.Append([]string{strconv.Itoa(r.Year), strconv.Itoa(r.Month), fmt.Sprintf("%.2f", r.Revenue)})
	}
	table.Render()

	fmt.Println()
}

func totalPurchaseMenu(reportService *service.ReportService) {
	fmt.Println()
	fmt.Println("=====================================")
	fmt.Println("Total Purchase")
	fmt.Println("=====================================")

	report, err := reportService.TotalPurchase()
	if err != nil {
		fmt.Println("Failed to get total purchase")
		return
	}

	if len(report) == 0 {
		fmt.Println("No data")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Product ID", "Product Name", "Count", "Total"})
	for _, r := range report {
		table.Append([]string{strconv.Itoa(r.ProductID), r.ProductName, strconv.Itoa(r.Count), strconv.FormatFloat(r.Total, 'f', 2, 64)})
	}
	table.Render()

	fmt.Println()
}

func totalUserMenu(reportService *service.ReportService) {
	fmt.Println()
	fmt.Println("=====================================")
	fmt.Println("Total User")
	fmt.Println("=====================================")

	report, err := reportService.TotalUser()
	if err != nil {
		fmt.Println("Failed to get total user")
		return
	}

	if len(report) == 0 {
		fmt.Println("No data")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Total"})
	for _, r := range report {
		table.Append([]string{strconv.Itoa(r.Total)})
	}
	table.Render()

	fmt.Println()
}
