// Code generated by mockery v1.0.0. DO NOT EDIT.

package emperror_test

import mock "github.com/stretchr/testify/mock"

// HandlerMock is an autogenerated mock type for the HandlerMock type
type HandlerMock struct {
	mock.Mock
}

// Handle provides a mock function with given fields: err
func (_m *HandlerMock) Handle(err error) {
	_m.Called(err)
}
