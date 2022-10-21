package users

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubStore struct {
	Data          []User
	Us            User
	ReadWasCalled bool
}

func (fs *StubStore) Read(data interface{}) error {
	fs.ReadWasCalled = true
	dataStub := data.(*[]User)
	*dataStub = fs.Data
	return nil
}

func (fs *StubStore) Write(data interface{}) error {
	// casterData := data.(User)
	// fs.Data = append(fs.Data, casterData)
	return nil
}

func TestGetAll(t *testing.T) {
	// Arrange
	us := []User{ //* Ej 1 con Stub
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

	// * Ej 2 con Mock
	userToUpdateMock := User{
		Id:       1,
		Name:     "Before Update",
		LastName: "Perez",
		Email:    "juan@gmail.com",
		Age:      39,
		Height:   1.78,
		Active:   true,
		Date:     "02-02-2022",
	}

	userExpected := User{
		Id:       1,
		Name:     "After Update",
		LastName: "Perez",
		Email:    "juan@gmail.com",
		Age:      39,
		Height:   1.78,
		Active:   true,
		Date:     "02-02-2022",
	}

	myStubStore := &StubStore{Data: us, Us: userToUpdateMock}
	repo := NewRepository(myStubStore)

	// Act
	result, err := repo.GetAll()
	resultUpdated, err := repo.UpdateName(userToUpdateMock.Id, userExpected.Name)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, us, result, "Deben ser iguales")
	assert.True(t, myStubStore.ReadWasCalled)
	assert.Equal(t, userExpected, resultUpdated, "Deben ser iguales")
}
