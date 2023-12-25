package types

import (
	prices "github.com/jinxprotocol/v4-chain/protocol/x/prices/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type PricesKeeper interface {
	GetMarketPrice(
		ctx sdk.Context,
		id uint32,
	) (market prices.MarketPrice, err error)
	// Methods imported from prices should be defined here
}
