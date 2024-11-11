// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	context "context"

	functions "github.com/smartcontractkit/chainlink/v2/core/services/functions"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// ORM is an autogenerated mock type for the ORM type
type ORM struct {
	mock.Mock
}

// CreateRequest provides a mock function with given fields: ctx, request
func (_m *ORM) CreateRequest(ctx context.Context, request *functions.Request) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for CreateRequest")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *functions.Request) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindById provides a mock function with given fields: ctx, requestID
func (_m *ORM) FindById(ctx context.Context, requestID functions.RequestID) (*functions.Request, error) {
	ret := _m.Called(ctx, requestID)

	if len(ret) == 0 {
		panic("no return value specified for FindById")
	}

	var r0 *functions.Request
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, functions.RequestID) (*functions.Request, error)); ok {
		return rf(ctx, requestID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, functions.RequestID) *functions.Request); ok {
		r0 = rf(ctx, requestID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*functions.Request)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, functions.RequestID) error); ok {
		r1 = rf(ctx, requestID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOldestEntriesByState provides a mock function with given fields: ctx, state, limit
func (_m *ORM) FindOldestEntriesByState(ctx context.Context, state functions.RequestState, limit uint32) ([]functions.Request, error) {
	ret := _m.Called(ctx, state, limit)

	if len(ret) == 0 {
		panic("no return value specified for FindOldestEntriesByState")
	}

	var r0 []functions.Request
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, functions.RequestState, uint32) ([]functions.Request, error)); ok {
		return rf(ctx, state, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, functions.RequestState, uint32) []functions.Request); ok {
		r0 = rf(ctx, state, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]functions.Request)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, functions.RequestState, uint32) error); ok {
		r1 = rf(ctx, state, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PruneOldestRequests provides a mock function with given fields: ctx, maxRequestsInDB, batchSize
func (_m *ORM) PruneOldestRequests(ctx context.Context, maxRequestsInDB uint32, batchSize uint32) (uint32, uint32, error) {
	ret := _m.Called(ctx, maxRequestsInDB, batchSize)

	if len(ret) == 0 {
		panic("no return value specified for PruneOldestRequests")
	}

	var r0 uint32
	var r1 uint32
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, uint32) (uint32, uint32, error)); ok {
		return rf(ctx, maxRequestsInDB, batchSize)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32, uint32) uint32); ok {
		r0 = rf(ctx, maxRequestsInDB, batchSize)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32, uint32) uint32); ok {
		r1 = rf(ctx, maxRequestsInDB, batchSize)
	} else {
		r1 = ret.Get(1).(uint32)
	}

	if rf, ok := ret.Get(2).(func(context.Context, uint32, uint32) error); ok {
		r2 = rf(ctx, maxRequestsInDB, batchSize)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SetConfirmed provides a mock function with given fields: ctx, requestID
func (_m *ORM) SetConfirmed(ctx context.Context, requestID functions.RequestID) error {
	ret := _m.Called(ctx, requestID)

	if len(ret) == 0 {
		panic("no return value specified for SetConfirmed")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, functions.RequestID) error); ok {
		r0 = rf(ctx, requestID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetError provides a mock function with given fields: ctx, requestID, errorType, computationError, readyAt, readyForProcessing
func (_m *ORM) SetError(ctx context.Context, requestID functions.RequestID, errorType functions.ErrType, computationError []byte, readyAt time.Time, readyForProcessing bool) error {
	ret := _m.Called(ctx, requestID, errorType, computationError, readyAt, readyForProcessing)

	if len(ret) == 0 {
		panic("no return value specified for SetError")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, functions.RequestID, functions.ErrType, []byte, time.Time, bool) error); ok {
		r0 = rf(ctx, requestID, errorType, computationError, readyAt, readyForProcessing)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetFinalized provides a mock function with given fields: ctx, requestID, reportedResult, reportedError
func (_m *ORM) SetFinalized(ctx context.Context, requestID functions.RequestID, reportedResult []byte, reportedError []byte) error {
	ret := _m.Called(ctx, requestID, reportedResult, reportedError)

	if len(ret) == 0 {
		panic("no return value specified for SetFinalized")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, functions.RequestID, []byte, []byte) error); ok {
		r0 = rf(ctx, requestID, reportedResult, reportedError)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetResult provides a mock function with given fields: ctx, requestID, computationResult, readyAt
func (_m *ORM) SetResult(ctx context.Context, requestID functions.RequestID, computationResult []byte, readyAt time.Time) error {
	ret := _m.Called(ctx, requestID, computationResult, readyAt)

	if len(ret) == 0 {
		panic("no return value specified for SetResult")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, functions.RequestID, []byte, time.Time) error); ok {
		r0 = rf(ctx, requestID, computationResult, readyAt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TimeoutExpiredResults provides a mock function with given fields: ctx, cutoff, limit
func (_m *ORM) TimeoutExpiredResults(ctx context.Context, cutoff time.Time, limit uint32) ([]functions.RequestID, error) {
	ret := _m.Called(ctx, cutoff, limit)

	if len(ret) == 0 {
		panic("no return value specified for TimeoutExpiredResults")
	}

	var r0 []functions.RequestID
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, uint32) ([]functions.RequestID, error)); ok {
		return rf(ctx, cutoff, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, uint32) []functions.RequestID); ok {
		r0 = rf(ctx, cutoff, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]functions.RequestID)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, time.Time, uint32) error); ok {
		r1 = rf(ctx, cutoff, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
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
