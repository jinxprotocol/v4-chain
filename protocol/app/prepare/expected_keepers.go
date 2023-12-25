package prepare

import (
	"time"

	bridgetypes "github.com/jinxprotocol/v4-chain/protocol/x/bridge/types"
	clobtypes "github.com/jinxprotocol/v4-chain/protocol/x/clob/types"
	perpstypes "github.com/jinxprotocol/v4-chain/protocol/x/perpetuals/types"
	pricestypes "github.com/jinxprotocol/v4-chain/protocol/x/prices/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// PrepareClobKeeper defines the expected CLOB keeper used for `PrepareProposal`.
type PrepareClobKeeper interface {
	GetOperations(ctx sdk.Context) *clobtypes.MsgProposedOperations
}

// PreparePerpetualsKeeper defines the expected Perpetuals keeper used for `PrepareProposal`.
type PreparePerpetualsKeeper interface {
	GetAddPremiumVotes(ctx sdk.Context) *perpstypes.MsgAddPremiumVotes
}

// PreparePricesKeeper defines the expected Prices keeper used for `PrepareProposal`.
type PreparePricesKeeper interface {
	GetValidMarketPriceUpdates(ctx sdk.Context) *pricestypes.MsgUpdateMarketPrices
}

// PrepareBridgeKeeper defines the expected Bridge keeper used for `PrepareProposal`.
type PrepareBridgeKeeper interface {
	GetAcknowledgeBridges(ctx sdk.Context, blockTimestamp time.Time) *bridgetypes.MsgAcknowledgeBridges
}
