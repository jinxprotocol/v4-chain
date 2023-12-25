package stats

import (
	"github.com/jinxprotocol/v4-chain/protocol/x/stats/keeper"
	"github.com/jinxprotocol/v4-chain/protocol/x/stats/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the stat module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.InitializeForGenesis(ctx)

	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the stat module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{
		Params: k.GetParams(ctx),
	}
}
