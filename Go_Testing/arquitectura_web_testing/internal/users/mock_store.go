package users

import "github.com/gcapodicimeli/backpack-bcgow6-gianfranco-capodici/arquitectura_web_testing/internal/domain"

type MockStorage struct {
	DataMock    []domain.User
	ErrOnWrite  error
	ErrOnRead   error
	readInvoked bool
}

func (m *MockStorage) Read(data interface{}) (err error) {
	if m.ErrOnRead != nil {
		return m.ErrOnRead
	}

	m.readInvoked = true
	castedData := data.(*[]domain.User)
	*castedData = m.DataMock
	return nil
}

func (m *MockStorage) Write(data interface{}) (err error) {
	if m.ErrOnWrite != nil {
		return m.ErrOnWrite
	}

	castedData := data.([]domain.User)
	m.DataMock = append(m.DataMock, castedData[len(castedData)-1])
	return nil
}
