package users

import "github.com/gcapodicimeli/backpack-bcgow6-gianfranco-capodici/arquitectura_web_testing/internal/domain"

type MockStorage struct {
	dataMock   []domain.User
	errOnWrite error
	errOnRead  error
}

func (m *MockStorage) Read(data interface{}) (err error) {
	if m.errOnRead != nil {
		return m.errOnRead
	}

	castedData := data.(*[]domain.User)
	*castedData = m.dataMock
	return nil
}

func (m *MockStorage) Write(data interface{}) (err error) {
	if m.errOnWrite != nil {
		return m.errOnWrite
	}

	castedData := data.(*domain.User)
	m.dataMock = append(m.dataMock, *castedData)
	return nil
}
