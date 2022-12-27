// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// IUuidService is an autogenerated mock type for the IUuidService type
type IUuidService struct {
	mock.Mock
}

// Random provides a mock function with given fields:
func (_m *IUuidService) Random() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewIUuidService interface {
	mock.TestingT
	Cleanup(func())
}

// NewIUuidService creates a new instance of IUuidService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIUuidService(t mockConstructorTestingTNewIUuidService) *IUuidService {
	mock := &IUuidService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}