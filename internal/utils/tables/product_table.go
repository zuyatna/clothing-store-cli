package tables

import (
	"clothing-pair-project/internal/models"
	"github.com/olekukonko/tablewriter"
	"strconv"
)

type ProductTable struct {
	table *tablewriter.Table
}

func ProductsTablePresenter(table *tablewriter.Table) *ProductTable {
	table.SetHeader([]string{"ID", "Name", "Price", "Description", "Image", "Type"})
	table.SetRowLine(true)
	return &ProductTable{table}
}

func (p *ProductTable) DisplayProducts(products []models.Product) {
	p.table.ClearRows()
	for _, product := range products {
		p.table.Append([]string{
			strconv.Itoa(product.ProductID),
			product.Name,
			"Rp." + strconv.FormatFloat(product.Price, 'f', 2, 64),
			product.Description.String,
			product.Images.String,
			product.Type,
		})
	}
	p.table.Render()
}

func (p *ProductTable) DisplayProduct(product models.Product) {
	p.table.ClearRows()
	p.table.Append([]string{
		strconv.Itoa(product.ProductID),
		product.Name,
		"Rp." + strconv.FormatFloat(product.Price, 'f', 2, 64),
		product.Description.String,
		product.Images.String,
		product.Type,
	})
	p.table.Render()
}
