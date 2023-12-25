package keeper

import (
	"context"

	gometrics "github.com/armon/go-metrics"
	pricefeedmetrics "github.com/jinxprotocol/v4-chain/protocol/daemons/pricefeed/metrics"
	"github.com/jinxprotocol/v4-chain/protocol/lib/metrics"

	"github.com/cosmos/cosmos-sdk/telemetry"

	errorsmod "cosmossdk.io/errors"
	"github.com/jinxprotocol/v4-chain/protocol/x/prices/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func (k msgServer) UpdateMarketParam(
	goCtx context.Context,
	msg *types.MsgUpdateMarketParam,
) (
	response *types.MsgUpdateMarketParamResponse,
	err error,
) {
	// Increment the appropriate success/error counter when the function finishes.
	defer func() {
		success := metrics.Success
		if err != nil {
			success = metrics.Error
		}
		telemetry.IncrCounterWithLabels(
			[]string{types.ModuleName, metrics.UpdateMarketParam, success},
			1,
			[]gometrics.Label{pricefeedmetrics.GetLabelForMarketId(msg.MarketParam.Id)},
		)
	}()

	if !k.Keeper.HasAuthority(msg.Authority) {
		return nil, errorsmod.Wrapf(
			govtypes.ErrInvalidSigner,
			"invalid authority %s",
			msg.Authority,
		)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if _, err = k.Keeper.ModifyMarketParam(ctx, msg.MarketParam); err != nil {
		return nil, err
	}

	return &types.MsgUpdateMarketParamResponse{}, nil
}
