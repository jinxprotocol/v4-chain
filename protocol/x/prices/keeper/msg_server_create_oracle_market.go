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

func (k msgServer) CreateOracleMarket(
	goCtx context.Context,
	msg *types.MsgCreateOracleMarket,
) (
	response *types.MsgCreateOracleMarketResponse,
	err error,
) {
	// Increment the appropriate success/error counter when the function finishes.
	defer func() {
		success := metrics.Success
		if err != nil {
			success = metrics.Error
		}
		telemetry.IncrCounterWithLabels(
			[]string{types.ModuleName, metrics.CreateOracleMarket, success},
			1,
			[]gometrics.Label{pricefeedmetrics.GetLabelForMarketId(msg.Params.Id)},
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

	// Use zero oracle price to create the new market.
	// Note that valid oracle price updates cannot be zero (checked in MsgUpdateMarketPrices.ValidateBasic),
	// so a zero oracle price indicates that the oracle price has never been updated.
	zeroMarketPrice := types.MarketPrice{
		Id:       msg.Params.Id,
		Exponent: msg.Params.Exponent,
		Price:    0,
	}
	if _, err = k.Keeper.CreateMarket(ctx, msg.Params, zeroMarketPrice); err != nil {
		return nil, err
	}

	return &types.MsgCreateOracleMarketResponse{}, nil
}
