package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"

	"github.com/jinxprotocol/v4-chain/protocol/x/clob/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func (k msgServer) UpdateClobPair(
	goCtx context.Context,
	msg *types.MsgUpdateClobPair,
) (resp *types.MsgUpdateClobPairResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.Keeper.HasAuthority(msg.Authority) {
		return nil, errorsmod.Wrapf(
			govtypes.ErrInvalidSigner,
			"invalid authority %s",
			msg.Authority,
		)
	}

	if err := k.Keeper.UpdateClobPair(
		ctx,
		msg.GetClobPair(),
	); err != nil {
		return nil, err
	}

	return &types.MsgUpdateClobPairResponse{}, nil
}
