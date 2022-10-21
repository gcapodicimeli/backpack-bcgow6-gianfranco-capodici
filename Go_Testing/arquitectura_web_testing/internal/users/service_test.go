package users

import (
	"encoding/json"
	"testing"

	"github.com/gcapodicimeli/backpack-bcgow6-gianfranco-capodici/arquitectura_web_testing/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestServiceIntegrationGetAll(t *testing.T) {
	// Arrange
	database := []User{
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

	db, _ := json.Marshal(database)
	mockStorage := store.Mock{
		Data:  db,
		Error: nil,
	}

	stubStore := store.FileStore{
		FilePath: "",
		Mock:     &mockStorage,
	}

	repository := NewRepository(&stubStore)
	service := NewService(repository)

	// Act
	result, err := service.GetAll()

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, database, result)
}
