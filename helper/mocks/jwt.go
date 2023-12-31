package mocks

import (
	jwt "github.com/golang-jwt/jwt/v5"
	mock "github.com/stretchr/testify/mock"
)

type JWTInterface struct {
	mock.Mock
}

func (_m *JWTInterface) ExtractToken(token *jwt.Token) interface{} {
	ret := _m.Called(token)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(*jwt.Token) interface{}); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

func (_m *JWTInterface) GenerateJWT(userID string) map[string]interface{} {
	ret := _m.Called(userID)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(string) map[string]interface{}); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	return r0
}

// GenerateToken provides a mock function with given fields: id
func (_m *JWTInterface) GenerateToken(id string) string {
	ret := _m.Called(id)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

func NewJWTInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *JWTInterface {
	mock := &JWTInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
