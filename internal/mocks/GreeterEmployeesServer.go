// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	employee "ChatsService/proto/employee"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// GreeterEmployeesServer is an autogenerated mock type for the GreeterEmployeesServer type
type GreeterEmployeesServer struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *GreeterEmployeesServer) Create(_a0 context.Context, _a1 *employee.EmployeeCreateRequest) (*employee.EmployeeCreateResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *employee.EmployeeCreateResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *employee.EmployeeCreateRequest) (*employee.EmployeeCreateResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *employee.EmployeeCreateRequest) *employee.EmployeeCreateResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*employee.EmployeeCreateResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *employee.EmployeeCreateRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Search provides a mock function with given fields: _a0, _a1
func (_m *GreeterEmployeesServer) Search(_a0 context.Context, _a1 *employee.SearchRequest) (*employee.SearchResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Search")
	}

	var r0 *employee.SearchResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *employee.SearchRequest) (*employee.SearchResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *employee.SearchRequest) *employee.SearchResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*employee.SearchResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *employee.SearchRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mustEmbedUnimplementedGreeterEmployeesServer provides a mock function with no fields
func (_m *GreeterEmployeesServer) mustEmbedUnimplementedGreeterEmployeesServer() {
	_m.Called()
}

// NewGreeterEmployeesServer creates a new instance of GreeterEmployeesServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGreeterEmployeesServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *GreeterEmployeesServer {
	mock := &GreeterEmployeesServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
