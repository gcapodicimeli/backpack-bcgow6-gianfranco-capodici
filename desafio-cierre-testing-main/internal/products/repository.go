package products

import "errors"

type Repository interface {
	GetAllBySeller(sellerID string) ([]Product, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAllBySeller(sellerID string) ([]Product, error) {
	if sellerID == "" {
		return []Product{}, errors.New("sellerID cannot be empty")
	}
	// Implementación para hacer fallar al 500 (que no sé si está bien)
	if sellerID == "#" {
		return []Product{}, errors.New("Invalid sellerID")
	}
	var prodList []Product
	prodList = append(prodList, Product{
		ID:          "mock",
		SellerID:    "FEX112AC",
		Description: "generic product",
		Price:       123.55,
	})
	return prodList, nil
}
