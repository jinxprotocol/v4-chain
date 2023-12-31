package keeper

import (
	"context"

	errorlib "github.com/jinxprotocol/v4-chain/protocol/lib/error"
	"github.com/jinxprotocol/v4-chain/protocol/lib/metrics"
	"github.com/jinxprotocol/v4-chain/protocol/x/clob/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ProposedOperations(
	goCtx context.Context,
	msg *types.MsgProposedOperations,
) (resp *types.MsgProposedOperationsResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	defer func() {
		metrics.IncrSuccessOrErrorCounter(err, types.ModuleName, metrics.ProposedOperations, metrics.DeliverTx)
		if err != nil {
			errorlib.LogDeliverTxError(k.Keeper.Logger(ctx), err, ctx.BlockHeight(), "ProposedOperations", msg)
		}
	}()

	if err := k.Keeper.ProcessProposerOperations(
		ctx,
		msg.GetOperationsQueue(),
	); err != nil {
		return nil, err
	}

	return &types.MsgProposedOperationsResponse{}, nil
}
