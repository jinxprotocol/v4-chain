package keeper

import (
	"context"

	"github.com/jinxprotocol/v4-chain/protocol/x/bridge/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AcknowledgeBridges acknowledges bridge events and sets them to complete
// at a later block.
func (k msgServer) AcknowledgeBridges(
	goCtx context.Context,
	msg *types.MsgAcknowledgeBridges,
) (*types.MsgAcknowledgeBridgesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.Keeper.AcknowledgeBridges(ctx, msg.Events); err != nil {
		return nil, err
	}

	return &types.MsgAcknowledgeBridgesResponse{}, nil
}
