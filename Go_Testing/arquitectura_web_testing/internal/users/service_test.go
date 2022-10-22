package users

import (
	"errors"
	"testing"

	"github.com/gcapodicimeli/backpack-bcgow6-gianfranco-capodici/arquitectura_web_testing/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestServiceIntegrationGetAll(t *testing.T) {
	// Arrange
	database := []domain.User{
		{Id: 1,
			Name:     "Juan",
			LastName: "Perez",
			Email:    "juan@gmail.com",
			Age:      39,
			Height:   1.78,
			Active:   true,
			Date:     "02-02-2022",
		},
		{Id: 2,
			Name:     "Juana",
			LastName: "Lopez",
			Email:    "juana@gmail.com",
			Age:      79,
			Height:   1.58,
			Active:   true,
			Date:     "02-06-2022",
		},
	}

	mockStorage := MockStorage{
		dataMock: database,
	}

	repository := NewRepository(&mockStorage)
	service := NewService(repository)

	// Act
	result, err := service.GetAll()

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, database, result)
}

func TestServiceIntegrationGetAllFail(t *testing.T) {
	// Arrange
	expectedError := errors.New("Mock error :)")
	mockStorage := MockStorage{
		dataMock:   nil,
		errOnWrite: nil,
		errOnRead:  errors.New("Mock error :)"),
	}

	repository := NewRepository(&mockStorage)
	service := NewService(repository)

	// Act
	result, err := service.GetAll()

	// Assert
	assert.EqualError(t, err, expectedError.Error())
	assert.Nil(t, result)
}
