package keeper

import (
	"context"

	"github.com/jinxprotocol/v4-chain/protocol/x/perpetuals/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) Params(
	c context.Context,
	req *types.QueryParamsRequest,
) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	params := k.GetParams(sdk.UnwrapSDKContext(c))

	return &types.QueryParamsResponse{Params: params}, nil
}
