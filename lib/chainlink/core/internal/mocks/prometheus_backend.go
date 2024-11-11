// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	big "math/big"

	mock "github.com/stretchr/testify/mock"
)

// PrometheusBackend is an autogenerated mock type for the PrometheusBackend type
type PrometheusBackend struct {
	mock.Mock
}

// SetMaxUnconfirmedAge provides a mock function with given fields: _a0, _a1
func (_m *PrometheusBackend) SetMaxUnconfirmedAge(_a0 *big.Int, _a1 float64) {
	_m.Called(_a0, _a1)
}

// SetMaxUnconfirmedBlocks provides a mock function with given fields: _a0, _a1
func (_m *PrometheusBackend) SetMaxUnconfirmedBlocks(_a0 *big.Int, _a1 int64) {
	_m.Called(_a0, _a1)
}

// SetPipelineRunsQueued provides a mock function with given fields: n
func (_m *PrometheusBackend) SetPipelineRunsQueued(n int) {
	_m.Called(n)
}

// SetPipelineTaskRunsQueued provides a mock function with given fields: n
func (_m *PrometheusBackend) SetPipelineTaskRunsQueued(n int) {
	_m.Called(n)
}

// SetUnconfirmedTransactions provides a mock function with given fields: _a0, _a1
func (_m *PrometheusBackend) SetUnconfirmedTransactions(_a0 *big.Int, _a1 int64) {
	_m.Called(_a0, _a1)
}

// NewPrometheusBackend creates a new instance of PrometheusBackend. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPrometheusBackend(t interface {
	mock.TestingT
	Cleanup(func())
}) *PrometheusBackend {
	mock := &PrometheusBackend{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
