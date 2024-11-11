// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	context "context"

	chaintype "github.com/smartcontractkit/chainlink/v2/core/services/keystore/chaintype"

	mock "github.com/stretchr/testify/mock"

	ocr2key "github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/ocr2key"
)

// OCR2 is an autogenerated mock type for the OCR2 type
type OCR2 struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, key
func (_m *OCR2) Add(ctx context.Context, key ocr2key.KeyBundle) error {
	ret := _m.Called(ctx, key)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ocr2key.KeyBundle) error); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *OCR2) Create(_a0 context.Context, _a1 chaintype.ChainType) (ocr2key.KeyBundle, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 ocr2key.KeyBundle
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, chaintype.ChainType) (ocr2key.KeyBundle, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, chaintype.ChainType) ocr2key.KeyBundle); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ocr2key.KeyBundle)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, chaintype.ChainType) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *OCR2) Delete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EnsureKeys provides a mock function with given fields: ctx, enabledChains
func (_m *OCR2) EnsureKeys(ctx context.Context, enabledChains ...chaintype.ChainType) error {
	_va := make([]interface{}, len(enabledChains))
	for _i := range enabledChains {
		_va[_i] = enabledChains[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for EnsureKeys")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...chaintype.ChainType) error); ok {
		r0 = rf(ctx, enabledChains...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Export provides a mock function with given fields: id, password
func (_m *OCR2) Export(id string, password string) ([]byte, error) {
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
func (_m *OCR2) Get(id string) (ocr2key.KeyBundle, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 ocr2key.KeyBundle
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (ocr2key.KeyBundle, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) ocr2key.KeyBundle); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ocr2key.KeyBundle)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *OCR2) GetAll() ([]ocr2key.KeyBundle, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []ocr2key.KeyBundle
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]ocr2key.KeyBundle, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []ocr2key.KeyBundle); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ocr2key.KeyBundle)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllOfType provides a mock function with given fields: _a0
func (_m *OCR2) GetAllOfType(_a0 chaintype.ChainType) ([]ocr2key.KeyBundle, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetAllOfType")
	}

	var r0 []ocr2key.KeyBundle
	var r1 error
	if rf, ok := ret.Get(0).(func(chaintype.ChainType) ([]ocr2key.KeyBundle, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(chaintype.ChainType) []ocr2key.KeyBundle); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ocr2key.KeyBundle)
		}
	}

	if rf, ok := ret.Get(1).(func(chaintype.ChainType) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Import provides a mock function with given fields: ctx, keyJSON, password
func (_m *OCR2) Import(ctx context.Context, keyJSON []byte, password string) (ocr2key.KeyBundle, error) {
	ret := _m.Called(ctx, keyJSON, password)

	if len(ret) == 0 {
		panic("no return value specified for Import")
	}

	var r0 ocr2key.KeyBundle
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []byte, string) (ocr2key.KeyBundle, error)); ok {
		return rf(ctx, keyJSON, password)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []byte, string) ocr2key.KeyBundle); ok {
		r0 = rf(ctx, keyJSON, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ocr2key.KeyBundle)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []byte, string) error); ok {
		r1 = rf(ctx, keyJSON, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewOCR2 creates a new instance of OCR2. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOCR2(t interface {
	mock.TestingT
	Cleanup(func())
}) *OCR2 {
	mock := &OCR2{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
