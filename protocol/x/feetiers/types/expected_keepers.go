package types

import (
	"github.com/jinxprotocol/v4-chain/protocol/x/stats/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// StatsKeeper defines the expected stats keeper
type StatsKeeper interface {
	GetUserStats(ctx sdk.Context, address string) *types.UserStats
	GetGlobalStats(ctx sdk.Context) *types.GlobalStats
}
