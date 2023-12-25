package keeper

import (
	"context"

	"github.com/jinxprotocol/v4-chain/protocol/x/vest/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) VestEntry(
	goCtx context.Context,
	req *types.QueryVestEntryRequest,
) (*types.QueryVestEntryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	vestEntry, err := k.GetVestEntry(ctx, req.VesterAccount)
	if err != nil {
		return nil, err
	}

	return &types.QueryVestEntryResponse{Entry: vestEntry}, nil
}
