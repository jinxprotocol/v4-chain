// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	daemonstypes "github.com/jinxprotocol/v4-chain/protocol/daemons/types"

	mock "github.com/stretchr/testify/mock"

	time "time"

	types "github.com/jinxprotocol/v4-chain/protocol/daemons/pricefeed/client/types"
)

// ExchangeQueryHandler is an autogenerated mock type for the ExchangeQueryHandler type
type ExchangeQueryHandler struct {
	mock.Mock
}

// Now provides a mock function with given fields:
func (_m *ExchangeQueryHandler) Now() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// Query provides a mock function with given fields: ctx, exchangeQueryDetails, exchangeConfig, marketIds, requestHandler, marketPriceExponent
func (_m *ExchangeQueryHandler) Query(ctx context.Context, exchangeQueryDetails *types.ExchangeQueryDetails, exchangeConfig *types.MutableExchangeMarketConfig, marketIds []uint32, requestHandler daemonstypes.RequestHandler, marketPriceExponent map[uint32]int32) ([]*types.MarketPriceTimestamp, map[uint32]error, error) {
	ret := _m.Called(ctx, exchangeQueryDetails, exchangeConfig, marketIds, requestHandler, marketPriceExponent)

	var r0 []*types.MarketPriceTimestamp
	if rf, ok := ret.Get(0).(func(context.Context, *types.ExchangeQueryDetails, *types.MutableExchangeMarketConfig, []uint32, daemonstypes.RequestHandler, map[uint32]int32) []*types.MarketPriceTimestamp); ok {
		r0 = rf(ctx, exchangeQueryDetails, exchangeConfig, marketIds, requestHandler, marketPriceExponent)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.MarketPriceTimestamp)
		}
	}

	var r1 map[uint32]error
	if rf, ok := ret.Get(1).(func(context.Context, *types.ExchangeQueryDetails, *types.MutableExchangeMarketConfig, []uint32, daemonstypes.RequestHandler, map[uint32]int32) map[uint32]error); ok {
		r1 = rf(ctx, exchangeQueryDetails, exchangeConfig, marketIds, requestHandler, marketPriceExponent)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[uint32]error)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *types.ExchangeQueryDetails, *types.MutableExchangeMarketConfig, []uint32, daemonstypes.RequestHandler, map[uint32]int32) error); ok {
		r2 = rf(ctx, exchangeQueryDetails, exchangeConfig, marketIds, requestHandler, marketPriceExponent)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type mockConstructorTestingTNewExchangeQueryHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewExchangeQueryHandler creates a new instance of ExchangeQueryHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewExchangeQueryHandler(t mockConstructorTestingTNewExchangeQueryHandler) *ExchangeQueryHandler {
	mock := &ExchangeQueryHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
