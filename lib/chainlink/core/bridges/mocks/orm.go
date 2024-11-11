// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	auth "github.com/smartcontractkit/chainlink/v2/core/auth"
	bridges "github.com/smartcontractkit/chainlink/v2/core/bridges"

	context "context"

	mock "github.com/stretchr/testify/mock"

	sqlutil "github.com/smartcontractkit/chainlink-common/pkg/sqlutil"

	time "time"
)

// ORM is an autogenerated mock type for the ORM type
type ORM struct {
	mock.Mock
}

// BridgeTypes provides a mock function with given fields: ctx, offset, limit
func (_m *ORM) BridgeTypes(ctx context.Context, offset int, limit int) ([]bridges.BridgeType, int, error) {
	ret := _m.Called(ctx, offset, limit)

	if len(ret) == 0 {
		panic("no return value specified for BridgeTypes")
	}

	var r0 []bridges.BridgeType
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) ([]bridges.BridgeType, int, error)); ok {
		return rf(ctx, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int) []bridges.BridgeType); ok {
		r0 = rf(ctx, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]bridges.BridgeType)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, int) int); ok {
		r1 = rf(ctx, offset, limit)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, int, int) error); ok {
		r2 = rf(ctx, offset, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// BulkUpsertBridgeResponse provides a mock function with given fields: ctx, responses
func (_m *ORM) BulkUpsertBridgeResponse(ctx context.Context, responses []bridges.BridgeResponse) error {
	ret := _m.Called(ctx, responses)

	if len(ret) == 0 {
		panic("no return value specified for BulkUpsertBridgeResponse")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []bridges.BridgeResponse) error); ok {
		r0 = rf(ctx, responses)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateBridgeType provides a mock function with given fields: ctx, bt
func (_m *ORM) CreateBridgeType(ctx context.Context, bt *bridges.BridgeType) error {
	ret := _m.Called(ctx, bt)

	if len(ret) == 0 {
		panic("no return value specified for CreateBridgeType")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *bridges.BridgeType) error); ok {
		r0 = rf(ctx, bt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateExternalInitiator provides a mock function with given fields: ctx, externalInitiator
func (_m *ORM) CreateExternalInitiator(ctx context.Context, externalInitiator *bridges.ExternalInitiator) error {
	ret := _m.Called(ctx, externalInitiator)

	if len(ret) == 0 {
		panic("no return value specified for CreateExternalInitiator")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *bridges.ExternalInitiator) error); ok {
		r0 = rf(ctx, externalInitiator)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteBridgeType provides a mock function with given fields: ctx, bt
func (_m *ORM) DeleteBridgeType(ctx context.Context, bt *bridges.BridgeType) error {
	ret := _m.Called(ctx, bt)

	if len(ret) == 0 {
		panic("no return value specified for DeleteBridgeType")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *bridges.BridgeType) error); ok {
		r0 = rf(ctx, bt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteExternalInitiator provides a mock function with given fields: ctx, name
func (_m *ORM) DeleteExternalInitiator(ctx context.Context, name string) error {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for DeleteExternalInitiator")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExternalInitiators provides a mock function with given fields: ctx, offset, limit
func (_m *ORM) ExternalInitiators(ctx context.Context, offset int, limit int) ([]bridges.ExternalInitiator, int, error) {
	ret := _m.Called(ctx, offset, limit)

	if len(ret) == 0 {
		panic("no return value specified for ExternalInitiators")
	}

	var r0 []bridges.ExternalInitiator
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) ([]bridges.ExternalInitiator, int, error)); ok {
		return rf(ctx, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int) []bridges.ExternalInitiator); ok {
		r0 = rf(ctx, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]bridges.ExternalInitiator)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, int) int); ok {
		r1 = rf(ctx, offset, limit)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, int, int) error); ok {
		r2 = rf(ctx, offset, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FindBridge provides a mock function with given fields: ctx, name
func (_m *ORM) FindBridge(ctx context.Context, name bridges.BridgeName) (bridges.BridgeType, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for FindBridge")
	}

	var r0 bridges.BridgeType
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, bridges.BridgeName) (bridges.BridgeType, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, bridges.BridgeName) bridges.BridgeType); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(bridges.BridgeType)
	}

	if rf, ok := ret.Get(1).(func(context.Context, bridges.BridgeName) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindBridges provides a mock function with given fields: ctx, name
func (_m *ORM) FindBridges(ctx context.Context, name []bridges.BridgeName) ([]bridges.BridgeType, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for FindBridges")
	}

	var r0 []bridges.BridgeType
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []bridges.BridgeName) ([]bridges.BridgeType, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []bridges.BridgeName) []bridges.BridgeType); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]bridges.BridgeType)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []bridges.BridgeName) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindExternalInitiator provides a mock function with given fields: ctx, eia
func (_m *ORM) FindExternalInitiator(ctx context.Context, eia *auth.Token) (*bridges.ExternalInitiator, error) {
	ret := _m.Called(ctx, eia)

	if len(ret) == 0 {
		panic("no return value specified for FindExternalInitiator")
	}

	var r0 *bridges.ExternalInitiator
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *auth.Token) (*bridges.ExternalInitiator, error)); ok {
		return rf(ctx, eia)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *auth.Token) *bridges.ExternalInitiator); ok {
		r0 = rf(ctx, eia)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bridges.ExternalInitiator)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *auth.Token) error); ok {
		r1 = rf(ctx, eia)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindExternalInitiatorByName provides a mock function with given fields: ctx, iname
func (_m *ORM) FindExternalInitiatorByName(ctx context.Context, iname string) (bridges.ExternalInitiator, error) {
	ret := _m.Called(ctx, iname)

	if len(ret) == 0 {
		panic("no return value specified for FindExternalInitiatorByName")
	}

	var r0 bridges.ExternalInitiator
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (bridges.ExternalInitiator, error)); ok {
		return rf(ctx, iname)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) bridges.ExternalInitiator); ok {
		r0 = rf(ctx, iname)
	} else {
		r0 = ret.Get(0).(bridges.ExternalInitiator)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, iname)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCachedResponse provides a mock function with given fields: ctx, dotId, specId, maxElapsed
func (_m *ORM) GetCachedResponse(ctx context.Context, dotId string, specId int32, maxElapsed time.Duration) ([]byte, error) {
	ret := _m.Called(ctx, dotId, specId, maxElapsed)

	if len(ret) == 0 {
		panic("no return value specified for GetCachedResponse")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int32, time.Duration) ([]byte, error)); ok {
		return rf(ctx, dotId, specId, maxElapsed)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int32, time.Duration) []byte); ok {
		r0 = rf(ctx, dotId, specId, maxElapsed)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int32, time.Duration) error); ok {
		r1 = rf(ctx, dotId, specId, maxElapsed)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCachedResponseWithFinished provides a mock function with given fields: ctx, dotId, specId, maxElapsed
func (_m *ORM) GetCachedResponseWithFinished(ctx context.Context, dotId string, specId int32, maxElapsed time.Duration) ([]byte, time.Time, error) {
	ret := _m.Called(ctx, dotId, specId, maxElapsed)

	if len(ret) == 0 {
		panic("no return value specified for GetCachedResponseWithFinished")
	}

	var r0 []byte
	var r1 time.Time
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int32, time.Duration) ([]byte, time.Time, error)); ok {
		return rf(ctx, dotId, specId, maxElapsed)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int32, time.Duration) []byte); ok {
		r0 = rf(ctx, dotId, specId, maxElapsed)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int32, time.Duration) time.Time); ok {
		r1 = rf(ctx, dotId, specId, maxElapsed)
	} else {
		r1 = ret.Get(1).(time.Time)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, int32, time.Duration) error); ok {
		r2 = rf(ctx, dotId, specId, maxElapsed)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateBridgeType provides a mock function with given fields: ctx, bt, btr
func (_m *ORM) UpdateBridgeType(ctx context.Context, bt *bridges.BridgeType, btr *bridges.BridgeTypeRequest) error {
	ret := _m.Called(ctx, bt, btr)

	if len(ret) == 0 {
		panic("no return value specified for UpdateBridgeType")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *bridges.BridgeType, *bridges.BridgeTypeRequest) error); ok {
		r0 = rf(ctx, bt, btr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpsertBridgeResponse provides a mock function with given fields: ctx, dotId, specId, response
func (_m *ORM) UpsertBridgeResponse(ctx context.Context, dotId string, specId int32, response []byte) error {
	ret := _m.Called(ctx, dotId, specId, response)

	if len(ret) == 0 {
		panic("no return value specified for UpsertBridgeResponse")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int32, []byte) error); ok {
		r0 = rf(ctx, dotId, specId, response)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WithDataSource provides a mock function with given fields: _a0
func (_m *ORM) WithDataSource(_a0 sqlutil.DataSource) bridges.ORM {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for WithDataSource")
	}

	var r0 bridges.ORM
	if rf, ok := ret.Get(0).(func(sqlutil.DataSource) bridges.ORM); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(bridges.ORM)
		}
	}

	return r0
}

// NewORM creates a new instance of ORM. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewORM(t interface {
	mock.TestingT
	Cleanup(func())
}) *ORM {
	mock := &ORM{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
