// Code generated by mockery. DO NOT EDIT.

// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023 Datadog, Inc.
package ddmark

import (
	reflect "reflect"

	mock "github.com/stretchr/testify/mock"
)

// DDValidationMarkerMock is an autogenerated mock type for the DDValidationMarker type
type DDValidationMarkerMock struct {
	mock.Mock
}

type DDValidationMarkerMock_Expecter struct {
	mock *mock.Mock
}

func (_m *DDValidationMarkerMock) EXPECT() *DDValidationMarkerMock_Expecter {
	return &DDValidationMarkerMock_Expecter{mock: &_m.Mock}
}

// ApplyRule provides a mock function with given fields: _a0
func (_m *DDValidationMarkerMock) ApplyRule(_a0 reflect.Value) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(reflect.Value) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DDValidationMarkerMock_ApplyRule_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ApplyRule'
type DDValidationMarkerMock_ApplyRule_Call struct {
	*mock.Call
}

// ApplyRule is a helper method to define mock.On call
//   - _a0 reflect.Value
func (_e *DDValidationMarkerMock_Expecter) ApplyRule(_a0 interface{}) *DDValidationMarkerMock_ApplyRule_Call {
	return &DDValidationMarkerMock_ApplyRule_Call{Call: _e.mock.On("ApplyRule", _a0)}
}

func (_c *DDValidationMarkerMock_ApplyRule_Call) Run(run func(_a0 reflect.Value)) *DDValidationMarkerMock_ApplyRule_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(reflect.Value))
	})
	return _c
}

func (_c *DDValidationMarkerMock_ApplyRule_Call) Return(_a0 error) *DDValidationMarkerMock_ApplyRule_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DDValidationMarkerMock_ApplyRule_Call) RunAndReturn(run func(reflect.Value) error) *DDValidationMarkerMock_ApplyRule_Call {
	_c.Call.Return(run)
	return _c
}

// TypeCheckError provides a mock function with given fields: _a0
func (_m *DDValidationMarkerMock) TypeCheckError(_a0 reflect.Value) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(reflect.Value) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DDValidationMarkerMock_TypeCheckError_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TypeCheckError'
type DDValidationMarkerMock_TypeCheckError_Call struct {
	*mock.Call
}

// TypeCheckError is a helper method to define mock.On call
//   - _a0 reflect.Value
func (_e *DDValidationMarkerMock_Expecter) TypeCheckError(_a0 interface{}) *DDValidationMarkerMock_TypeCheckError_Call {
	return &DDValidationMarkerMock_TypeCheckError_Call{Call: _e.mock.On("TypeCheckError", _a0)}
}

func (_c *DDValidationMarkerMock_TypeCheckError_Call) Run(run func(_a0 reflect.Value)) *DDValidationMarkerMock_TypeCheckError_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(reflect.Value))
	})
	return _c
}

func (_c *DDValidationMarkerMock_TypeCheckError_Call) Return(_a0 error) *DDValidationMarkerMock_TypeCheckError_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DDValidationMarkerMock_TypeCheckError_Call) RunAndReturn(run func(reflect.Value) error) *DDValidationMarkerMock_TypeCheckError_Call {
	_c.Call.Return(run)
	return _c
}

// ValueCheckError provides a mock function with given fields:
func (_m *DDValidationMarkerMock) ValueCheckError() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DDValidationMarkerMock_ValueCheckError_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ValueCheckError'
type DDValidationMarkerMock_ValueCheckError_Call struct {
	*mock.Call
}

// ValueCheckError is a helper method to define mock.On call
func (_e *DDValidationMarkerMock_Expecter) ValueCheckError() *DDValidationMarkerMock_ValueCheckError_Call {
	return &DDValidationMarkerMock_ValueCheckError_Call{Call: _e.mock.On("ValueCheckError")}
}

func (_c *DDValidationMarkerMock_ValueCheckError_Call) Run(run func()) *DDValidationMarkerMock_ValueCheckError_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DDValidationMarkerMock_ValueCheckError_Call) Return(_a0 error) *DDValidationMarkerMock_ValueCheckError_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DDValidationMarkerMock_ValueCheckError_Call) RunAndReturn(run func() error) *DDValidationMarkerMock_ValueCheckError_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewDDValidationMarkerMock interface {
	mock.TestingT
	Cleanup(func())
}

// NewDDValidationMarkerMock creates a new instance of DDValidationMarkerMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDDValidationMarkerMock(t mockConstructorTestingTNewDDValidationMarkerMock) *DDValidationMarkerMock {
	mock := &DDValidationMarkerMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
