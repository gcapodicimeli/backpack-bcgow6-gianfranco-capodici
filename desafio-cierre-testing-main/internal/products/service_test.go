package products

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ?- Esto es test unitario o de integración?
func TestGetAllBySeller(t *testing.T) {
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

func TestGetAllBySellerFail(t *testing.T) {
	// Arrange
	sellerID := ""
	expectedError := errors.New("sellerID cannot be empty")
	repository := NewRepository()
	service := NewService(repository)

	// Act
	_, err := service.GetAllBySeller(sellerID)

	// Assert
	assert.EqualError(t, expectedError, err.Error())
}
