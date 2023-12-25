package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"

	"github.com/jinxprotocol/v4-chain/protocol/x/clob/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// UpdateLiquidationsConfig updates the liquidation config in state.
func (k msgServer) UpdateLiquidationsConfig(
	goCtx context.Context,
	msg *types.MsgUpdateLiquidationsConfig,
) (resp *types.MsgUpdateLiquidationsConfigResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.Keeper.HasAuthority(msg.Authority) {
		return nil, errorsmod.Wrapf(
			govtypes.ErrInvalidSigner,
			"invalid authority %s",
			msg.Authority,
		)
	}

	if err := k.Keeper.UpdateLiquidationsConfig(ctx, msg.LiquidationsConfig); err != nil {
		return nil, err
	}
	return &types.MsgUpdateLiquidationsConfigResponse{}, nil
}
