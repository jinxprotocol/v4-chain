package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/jinxprotocol/v4-chain/protocol/x/perpetuals/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func (k msgServer) SetLiquidityTier(
	goCtx context.Context,
	msg *types.MsgSetLiquidityTier,
) (*types.MsgSetLiquidityTierResponse, error) {
	if !k.Keeper.HasAuthority(msg.Authority) {
		return nil, errorsmod.Wrapf(
			govtypes.ErrInvalidSigner,
			"invalid authority %s",
			msg.Authority,
		)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if _, err := k.Keeper.SetLiquidityTier(
		ctx,
		msg.LiquidityTier.Id,
		msg.LiquidityTier.Name,
		msg.LiquidityTier.InitialMarginPpm,
		msg.LiquidityTier.MaintenanceFractionPpm,
		msg.LiquidityTier.ImpactNotional,
	); err != nil {
		return nil, err
	}

	return &types.MsgSetLiquidityTierResponse{}, nil
}
