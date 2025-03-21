// Code generated by mockery. DO NOT EDIT.

// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023 Datadog, Inc.
package eventnotifier

import (
	types "github.com/DataDog/chaos-controller/eventnotifier/types"
	mock "github.com/stretchr/testify/mock"
	v1 "k8s.io/api/core/v1"

	v1beta1 "github.com/DataDog/chaos-controller/api/v1beta1"
)

// NotifierMock is an autogenerated mock type for the Notifier type
type NotifierMock struct {
	mock.Mock
}

type NotifierMock_Expecter struct {
	mock *mock.Mock
}

func (_m *NotifierMock) EXPECT() *NotifierMock_Expecter {
	return &NotifierMock_Expecter{mock: &_m.Mock}
}

// GetNotifierName provides a mock function with given fields:
func (_m *NotifierMock) GetNotifierName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NotifierMock_GetNotifierName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetNotifierName'
type NotifierMock_GetNotifierName_Call struct {
	*mock.Call
}

// GetNotifierName is a helper method to define mock.On call
func (_e *NotifierMock_Expecter) GetNotifierName() *NotifierMock_GetNotifierName_Call {
	return &NotifierMock_GetNotifierName_Call{Call: _e.mock.On("GetNotifierName")}
}

func (_c *NotifierMock_GetNotifierName_Call) Run(run func()) *NotifierMock_GetNotifierName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *NotifierMock_GetNotifierName_Call) Return(_a0 string) *NotifierMock_GetNotifierName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *NotifierMock_GetNotifierName_Call) RunAndReturn(run func() string) *NotifierMock_GetNotifierName_Call {
	_c.Call.Return(run)
	return _c
}

// Notify provides a mock function with given fields: _a0, _a1, _a2
func (_m *NotifierMock) Notify(_a0 v1beta1.Disruption, _a1 v1.Event, _a2 types.NotificationType) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(v1beta1.Disruption, v1.Event, types.NotificationType) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NotifierMock_Notify_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Notify'
type NotifierMock_Notify_Call struct {
	*mock.Call
}

// Notify is a helper method to define mock.On call
//   - _a0 v1beta1.Disruption
//   - _a1 v1.Event
//   - _a2 types.NotificationType
func (_e *NotifierMock_Expecter) Notify(_a0 interface{}, _a1 interface{}, _a2 interface{}) *NotifierMock_Notify_Call {
	return &NotifierMock_Notify_Call{Call: _e.mock.On("Notify", _a0, _a1, _a2)}
}

func (_c *NotifierMock_Notify_Call) Run(run func(_a0 v1beta1.Disruption, _a1 v1.Event, _a2 types.NotificationType)) *NotifierMock_Notify_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(v1beta1.Disruption), args[1].(v1.Event), args[2].(types.NotificationType))
	})
	return _c
}

func (_c *NotifierMock_Notify_Call) Return(_a0 error) *NotifierMock_Notify_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *NotifierMock_Notify_Call) RunAndReturn(run func(v1beta1.Disruption, v1.Event, types.NotificationType) error) *NotifierMock_Notify_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewNotifierMock interface {
	mock.TestingT
	Cleanup(func())
}

// NewNotifierMock creates a new instance of NotifierMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNotifierMock(t mockConstructorTestingTNewNotifierMock) *NotifierMock {
	mock := &NotifierMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
