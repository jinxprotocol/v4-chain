package types

import (
	"github.com/jinxprotocol/v4-chain/protocol/x/epochs/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// EpochsKeeper defines the expected epochs keeper to get epoch info.
type EpochsKeeper interface {
	MustGetStatsEpochInfo(ctx sdk.Context) types.EpochInfo
}
