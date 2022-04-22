package mocks

import (
	mock "github.com/stretchr/testify/mock"
)

type Fallback struct {
	mock.Mock
}

func (_m *Fallback) Method(topic string, payload interface{}) {
	_m.Called(topic, payload)
}
