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

// func TestServiceIntegrationStore(t *testing.T) {
// 	// Arrange
// 	initiaDatabase := []domain.User{
// 		{Id: 1,
// 			Name:     "Juan",
// 			LastName: "Perez",
// 			Email:    "juan@gmail.com",
// 			Age:      39,
// 			Height:   1.78,
// 			Active:   true,
// 			Date:     "02-02-2022",
// 		},
// 		{Id: 2,
// 			Name:     "Juana",
// 			LastName: "Lopez",
// 			Email:    "juana@gmail.com",
// 			Age:      79,
// 			Height:   1.58,
// 			Active:   true,
// 			Date:     "02-06-2022",
// 		},
// 	}

// 	expectedDatabase := []domain.User{
// 		{Id: 1,
// 			Name:     "Juan",
// 			LastName: "Perez",
// 			Email:    "juan@gmail.com",
// 			Age:      39,
// 			Height:   1.78,
// 			Active:   true,
// 			Date:     "02-02-2022",
// 		},
// 		{Id: 2,
// 			Name:     "Juana",
// 			LastName: "Lopez",
// 			Email:    "juana@gmail.com",
// 			Age:      79,
// 			Height:   1.58,
// 			Active:   true,
// 			Date:     "02-06-2022",
// 		},
// 		{
// 			Id:       3,
// 			Name:     "Maria",
// 			LastName: "Pereyra",
// 			Email:    "maria@gmail.com",
// 			Age:      45,
// 			Height:   1.65,
// 			Active:   true,
// 			Date:     "12-09-2022",
// 		},
// 	}

// 	mockStorage := MockStorage{
// 		dataMock: initiaDatabase,
// 	}

// 	repository := NewRepository(&mockStorage)
// 	service := NewService(repository)

// 	// Act
// 	result, err := service.Store("Maria", "Pereyra", "maria@gmail.com", 45, 1.65, true, "12-09-2022")
// 	userCreated := domain.User{
// 		Id:       3,
// 		Name:     "Maria",
// 		LastName: "Pereyra",
// 		Email:    "maria@gmail.com",
// 		Age:      45,
// 		Height:   1.65,
// 		Active:   true,
// 		Date:     "12-09-2022",
// 	}

// 	// Assert
// 	assert.Nil(t, err)
// 	assert.Equal(t, userCreated, result)
// 	assert.Equal(t, expectedDatabase, mockStorage.dataMock)
// }

func TestServiceIntegrationStoreFail(t *testing.T) {
	// Arrange
	expectedError := errors.New("Error on write a user")
	mockStorage := MockStorage{
		dataMock:   nil,
		errOnRead:  nil,
		errOnWrite: errors.New("Error on write a user"),
	}

	repository := NewRepository(&mockStorage)
	service := NewService(repository)

	// Act
	result, err := service.Store("Maria", "Perez", "maria@gmail.com", 45, 1.65, true, "12-03-2022")

	// Assert
	assert.EqualError(t, err, expectedError.Error())
	assert.Empty(t, result)
}
