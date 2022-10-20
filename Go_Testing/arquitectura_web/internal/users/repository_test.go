package users

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubStore struct {
	Data []User
}

func (fs *StubStore) Read(data interface{}) error {
	dataStub := data.(*[]User)
	*dataStub = fs.Data
	return nil
}

func TestGetAll(t *testing.T) {
	// Arrange
	us := []User{Id: 1,
		Name:     "Juan",
		LastName: "Perez",
		Email:    "juan@gmail.com",
		Age:      39,
		Height:   1.78,
		Active:   true,
		Data:     "02-02-2022",
	}
	myStubStore := &StubStore{Data: us}
	repo := NewRepository(myStubStore)

	// Act
	result, err := repo.GetAll()

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, us, result, "Deben ser iguales")
}
