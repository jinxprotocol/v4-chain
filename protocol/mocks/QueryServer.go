// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	types "github.com/jinxprotocol/v4-chain/protocol/x/prices/types"
)

// QueryServer is an autogenerated mock type for the QueryServer type
type QueryServer struct {
	mock.Mock
}

// AllMarketParams provides a mock function with given fields: _a0, _a1
func (_m *QueryServer) AllMarketParams(_a0 context.Context, _a1 *types.QueryAllMarketParamsRequest) (*types.QueryAllMarketParamsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.QueryAllMarketParamsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAllMarketParamsRequest) *types.QueryAllMarketParamsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryAllMarketParamsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryAllMarketParamsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AllMarketPrices provides a mock function with given fields: _a0, _a1
func (_m *QueryServer) AllMarketPrices(_a0 context.Context, _a1 *types.QueryAllMarketPricesRequest) (*types.QueryAllMarketPricesResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.QueryAllMarketPricesResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAllMarketPricesRequest) *types.QueryAllMarketPricesResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryAllMarketPricesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryAllMarketPricesRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MarketParam provides a mock function with given fields: _a0, _a1
func (_m *QueryServer) MarketParam(_a0 context.Context, _a1 *types.QueryMarketParamRequest) (*types.QueryMarketParamResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.QueryMarketParamResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryMarketParamRequest) *types.QueryMarketParamResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryMarketParamResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryMarketParamRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MarketPrice provides a mock function with given fields: _a0, _a1
func (_m *QueryServer) MarketPrice(_a0 context.Context, _a1 *types.QueryMarketPriceRequest) (*types.QueryMarketPriceResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.QueryMarketPriceResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryMarketPriceRequest) *types.QueryMarketPriceResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryMarketPriceResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryMarketPriceRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewQueryServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewQueryServer creates a new instance of QueryServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewQueryServer(t mockConstructorTestingTNewQueryServer) *QueryServer {
	mock := &QueryServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
