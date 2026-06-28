package calculator

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockDBService struct {
	mock.Mock
}

func (m *MockDBService) GetSum() int {
	args := m.Called()
	return args.Int(0)
}

func TestGetSumFromDB(t *testing.T) {
	// Create a mock
	mockDB := new(MockDBService)

	// Set the expectation for the GetSum call and return 10
	mockDB.On("GetSum").Return(10)

	// Call the function using the mock
	result := GetSumFromDB(mockDB)

	// Verify that the expectation was met
	mockDB.AssertExpectations(t)

	// Verify that the result is as expected
	if result != 10 {
		t.Errorf("Expected result to be 10, but got %d", result)
	}
}
