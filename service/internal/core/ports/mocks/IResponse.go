// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	fiber "github.com/gofiber/fiber/v2"
	mock "github.com/stretchr/testify/mock"
)

// IResponse is an autogenerated mock type for the IResponse type
type IResponse struct {
	mock.Mock
}

// ErrorRequestBody provides a mock function with given fields: h
func (_m *IResponse) ErrorRequestBody(h *fiber.Ctx) error {
	ret := _m.Called(h)

	var r0 error
	if rf, ok := ret.Get(0).(func(*fiber.Ctx) error); ok {
		r0 = rf(h)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Json provides a mock function with given fields: h, httpCode, message, result, status
func (_m *IResponse) Json(h *fiber.Ctx, httpCode int, message string, result interface{}, status bool) error {
	ret := _m.Called(h, httpCode, message, result, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(*fiber.Ctx, int, string, interface{}, bool) error); ok {
		r0 = rf(h, httpCode, message, result, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// JsonAuth provides a mock function with given fields: h, httpCode, message, result, status, isValidAuthorize
func (_m *IResponse) JsonAuth(h *fiber.Ctx, httpCode int, message string, result interface{}, status bool, isValidAuthorize bool) error {
	ret := _m.Called(h, httpCode, message, result, status, isValidAuthorize)

	var r0 error
	if rf, ok := ret.Get(0).(func(*fiber.Ctx, int, string, interface{}, bool, bool) error); ok {
		r0 = rf(h, httpCode, message, result, status, isValidAuthorize)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewIResponse interface {
	mock.TestingT
	Cleanup(func())
}

// NewIResponse creates a new instance of IResponse. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIResponse(t mockConstructorTestingTNewIResponse) *IResponse {
	mock := &IResponse{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
