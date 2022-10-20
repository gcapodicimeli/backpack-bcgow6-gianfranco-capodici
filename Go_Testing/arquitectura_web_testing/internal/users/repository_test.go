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

func (fs *StubStore) Write(data interface{}) error {
	casterData := data.(User)
	fs.Data = append(fs.Data, casterData)
	return nil
}

func TestGetAll(t *testing.T) {
	// Arrange
	us := []User{
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

	myStubStore := &StubStore{Data: us}
	repo := NewRepository(myStubStore)

	// Act
	result, err := repo.GetAll()

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, us, result, "Deben ser iguales")
}
