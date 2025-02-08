package handler

import (
	"clothing-pair-project/internal/utils/interfaces"
	"fmt"
)

type ProductHandler struct {
	productDisplay interfaces.ProductDisplay
	productFetcher interfaces.ProductFetcher
}

func NewProductHandler(productDisplay interfaces.ProductDisplay, productFetcher interfaces.ProductFetcher) *ProductHandler {
	return &ProductHandler{
		productDisplay: productDisplay,
		productFetcher: productFetcher,
	}
}

func (h *ProductHandler) ShowAllProducts(limit, offset int) (bool, bool, error) {
	products, err := h.productFetcher.GetAllProducts(limit+1, offset)
	if err != nil {
		return false, false, fmt.Errorf("error fetching all products: %w", err)
	}

	if len(products) == 0 {
		return false, false, fmt.Errorf("no products found")
	}

	displayProducts := products
	if len(products) > limit {
		displayProducts = products[:limit]
	}

	h.productDisplay.DisplayProducts(displayProducts)

	hasNext := len(products) > limit
	hasPrev := offset > 0

	return hasNext, hasPrev, nil
}

func (h *ProductHandler) ShowProductByID(productID int) error {
	product, err := h.productFetcher.GetProductByID(productID)
	if err != nil {
		return fmt.Errorf("error fetching product by ID: %w", err)
	}

	h.productDisplay.DisplayProduct(product)

	return nil
}
