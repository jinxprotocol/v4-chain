// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	cosmos_sdktypes "github.com/cosmos/cosmos-sdk/types"

	types "github.com/jinxprotocol/v4-chain/protocol/x/sending/types"
)

// SendingKeeper is an autogenerated mock type for the SendingKeeper type
type SendingKeeper struct {
	mock.Mock
}

// HasAuthority provides a mock function with given fields: authority
func (_m *SendingKeeper) HasAuthority(authority string) bool {
	ret := _m.Called(authority)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(authority)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ProcessDepositToSubaccount provides a mock function with given fields: ctx, msgDepositToSubaccount
func (_m *SendingKeeper) ProcessDepositToSubaccount(ctx cosmos_sdktypes.Context, msgDepositToSubaccount *types.MsgDepositToSubaccount) error {
	ret := _m.Called(ctx, msgDepositToSubaccount)

	var r0 error
	if rf, ok := ret.Get(0).(func(cosmos_sdktypes.Context, *types.MsgDepositToSubaccount) error); ok {
		r0 = rf(ctx, msgDepositToSubaccount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProcessTransfer provides a mock function with given fields: ctx, transfer
func (_m *SendingKeeper) ProcessTransfer(ctx cosmos_sdktypes.Context, transfer *types.Transfer) error {
	ret := _m.Called(ctx, transfer)

	var r0 error
	if rf, ok := ret.Get(0).(func(cosmos_sdktypes.Context, *types.Transfer) error); ok {
		r0 = rf(ctx, transfer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProcessWithdrawFromSubaccount provides a mock function with given fields: ctx, msgWithdrawFromSubaccount
func (_m *SendingKeeper) ProcessWithdrawFromSubaccount(ctx cosmos_sdktypes.Context, msgWithdrawFromSubaccount *types.MsgWithdrawFromSubaccount) error {
	ret := _m.Called(ctx, msgWithdrawFromSubaccount)

	var r0 error
	if rf, ok := ret.Get(0).(func(cosmos_sdktypes.Context, *types.MsgWithdrawFromSubaccount) error); ok {
		r0 = rf(ctx, msgWithdrawFromSubaccount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendFromModuleToAccount provides a mock function with given fields: ctx, msg
func (_m *SendingKeeper) SendFromModuleToAccount(ctx cosmos_sdktypes.Context, msg *types.MsgSendFromModuleToAccount) error {
	ret := _m.Called(ctx, msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(cosmos_sdktypes.Context, *types.MsgSendFromModuleToAccount) error); ok {
		r0 = rf(ctx, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewSendingKeeper interface {
	mock.TestingT
	Cleanup(func())
}

// NewSendingKeeper creates a new instance of SendingKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSendingKeeper(t mockConstructorTestingTNewSendingKeeper) *SendingKeeper {
	mock := &SendingKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
