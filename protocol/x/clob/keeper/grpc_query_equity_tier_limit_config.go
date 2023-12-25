package keeper

import (
	"context"

	"google.golang.org/grpc/status"

	"github.com/jinxprotocol/v4-chain/protocol/x/clob/types"
	"google.golang.org/grpc/codes"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) EquityTierLimitConfiguration(
	c context.Context,
	req *types.QueryEquityTierLimitConfigurationRequest,
) (*types.QueryEquityTierLimitConfigurationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryEquityTierLimitConfigurationResponse{
		EquityTierLimitConfig: k.GetEquityTierLimitConfiguration(ctx),
	}, nil
}
