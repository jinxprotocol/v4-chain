package epochs

import (
	"github.com/jinxprotocol/v4-chain/protocol/x/epochs/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BeginBlocker executes all ABCI BeginBlock logic respective to the epochs module.
func BeginBlocker(ctx sdk.Context, keeper keeper.Keeper) {
	epochs := keeper.GetAllEpochInfo(ctx)
	// Iterate through all epoch infos, calls MaybeStartNextEpoch() which
	// initializes and/or increments the epoch if applicable.
	for _, epoch := range epochs {
		if _, err := keeper.MaybeStartNextEpoch(ctx, epoch.GetEpochInfoName()); err != nil {
			panic(err)
		}
	}
}
