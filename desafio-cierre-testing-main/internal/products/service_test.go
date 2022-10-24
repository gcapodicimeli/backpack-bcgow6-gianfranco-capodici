package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllBySellerStub(t *testing.T) {
	// Arrange
	sellerID := "mock"

	repository := NewRepository()
	service := NewService(repository)

	expectedProduct := []Product{
		{
			ID:          "mock",
			SellerID:    "FEX112AC",
			Description: "generic product",
			Price:       123.55,
		},
	}

	// Act
	result, err := service.GetAllBySeller(sellerID)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, expectedProduct, result)
}
