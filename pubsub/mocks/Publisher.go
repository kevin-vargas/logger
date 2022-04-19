// Code generated by mockery v2.10.6. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Publisher is an autogenerated mock type for the Publisher type
type Publisher struct {
	mock.Mock
}

// Publish provides a mock function with given fields: topic, payload
func (_m *Publisher) Publish(topic string, payload interface{}) error {
	ret := _m.Called(topic, payload)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(topic, payload)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
