package mocks

import (
	"github.com/ivchip/go-meli-filter-ip/domain"
	"github.com/stretchr/testify/mock"
)

type MockContextualResultUseCases struct {
	mock.Mock
}

func (m *MockContextualResultUseCases) GetByIP(ip string) (domain.ContextualResult, error) {
	args := m.Called(ip)
	return args.Get(0).(domain.ContextualResult), args.Error(1)
}
