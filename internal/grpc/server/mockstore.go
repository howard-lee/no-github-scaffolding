package server

import (
	grpcApi "api"

	"github.com/stretchr/testify/mock"
)

// MockStore is a mock implementation of a datastore for testing purposes
type MockStore struct {
	mock.Mock
}

//Search returns the object which was passed to the mock on setup
func (m *MockStore) Search(name string) []grpcApi.Balance {
	args := m.Mock.Called(name)

	return args.Get(0).([]grpcApi.Balance)
}
