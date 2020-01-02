package mock

import (
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockCloser struct {
	mock.Mock
}

func (c *MockCloser) Close() {
	c.Called()
}

func Test_SecurityCentral_SecurityOn(t *testing.T) {
	var mockWindow MockCloser
	var mockDoor MockCloser

	mockWindow.On("Close").Return()
	mockDoor.On("Close").Return()

	sc := NewSecurityCentral(&mockWindow, &mockDoor)
	sc.SecurityOn()

	mockWindow.AssertExpectations(t)
	mockDoor.AssertExpectations(t)
}
