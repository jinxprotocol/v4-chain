package keeper

import (
	"context"

	"github.com/jinxprotocol/v4-chain/protocol/x/clob/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) BlockRateLimitConfiguration(
	c context.Context,
	req *types.QueryBlockRateLimitConfigurationRequest,
) (*types.QueryBlockRateLimitConfigurationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	blockRateLimitConfig := k.GetBlockRateLimitConfiguration(ctx)

	return &types.QueryBlockRateLimitConfigurationResponse{
		BlockRateLimitConfig: blockRateLimitConfig,
	}, nil
}
