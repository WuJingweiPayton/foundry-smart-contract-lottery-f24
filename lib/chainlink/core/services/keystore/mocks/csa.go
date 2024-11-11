// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	context "context"

	csakey "github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/csakey"

	mock "github.com/stretchr/testify/mock"
)

// CSA is an autogenerated mock type for the CSA type
type CSA struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, key
func (_m *CSA) Add(ctx context.Context, key csakey.KeyV2) error {
	ret := _m.Called(ctx, key)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, csakey.KeyV2) error); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: ctx
func (_m *CSA) Create(ctx context.Context) (csakey.KeyV2, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 csakey.KeyV2
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (csakey.KeyV2, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) csakey.KeyV2); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(csakey.KeyV2)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *CSA) Delete(ctx context.Context, id string) (csakey.KeyV2, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 csakey.KeyV2
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (csakey.KeyV2, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) csakey.KeyV2); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(csakey.KeyV2)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EnsureKey provides a mock function with given fields: ctx
func (_m *CSA) EnsureKey(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for EnsureKey")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Export provides a mock function with given fields: id, password
func (_m *CSA) Export(id string, password string) ([]byte, error) {
	ret := _m.Called(id, password)

	if len(ret) == 0 {
		panic("no return value specified for Export")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) ([]byte, error)); ok {
		return rf(id, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) []byte); ok {
		r0 = rf(id, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(id, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: id
func (_m *CSA) Get(id string) (csakey.KeyV2, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 csakey.KeyV2
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (csakey.KeyV2, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) csakey.KeyV2); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(csakey.KeyV2)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *CSA) GetAll() ([]csakey.KeyV2, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []csakey.KeyV2
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]csakey.KeyV2, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []csakey.KeyV2); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]csakey.KeyV2)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Import provides a mock function with given fields: ctx, keyJSON, password
func (_m *CSA) Import(ctx context.Context, keyJSON []byte, password string) (csakey.KeyV2, error) {
	ret := _m.Called(ctx, keyJSON, password)

	if len(ret) == 0 {
		panic("no return value specified for Import")
	}

	var r0 csakey.KeyV2
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []byte, string) (csakey.KeyV2, error)); ok {
		return rf(ctx, keyJSON, password)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []byte, string) csakey.KeyV2); ok {
		r0 = rf(ctx, keyJSON, password)
	} else {
		r0 = ret.Get(0).(csakey.KeyV2)
	}

	if rf, ok := ret.Get(1).(func(context.Context, []byte, string) error); ok {
		r1 = rf(ctx, keyJSON, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCSA creates a new instance of CSA. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCSA(t interface {
	mock.TestingT
	Cleanup(func())
}) *CSA {
	mock := &CSA{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
