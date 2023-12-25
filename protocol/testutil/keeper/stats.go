package keeper

import (
	"github.com/jinxprotocol/v4-chain/protocol/lib"
	"github.com/jinxprotocol/v4-chain/protocol/mocks"
	delaymsgtypes "github.com/jinxprotocol/v4-chain/protocol/x/delaymsg/types"

	tmdb "github.com/cometbft/cometbft-db"
	epochskeeper "github.com/jinxprotocol/v4-chain/protocol/x/epochs/keeper"
	"github.com/jinxprotocol/v4-chain/protocol/x/stats/keeper"
	"github.com/jinxprotocol/v4-chain/protocol/x/stats/types"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func createStatsKeeper(
	stateStore storetypes.CommitMultiStore,
	epochsKeeper *epochskeeper.Keeper,
	db *tmdb.MemDB,
	cdc *codec.ProtoCodec,
) (*keeper.Keeper, storetypes.StoreKey) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	transientStoreKey := sdk.NewTransientStoreKey(types.TransientStoreKey)

	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(transientStoreKey, storetypes.StoreTypeTransient, db)

	mockMsgSender := &mocks.IndexerMessageSender{}
	mockMsgSender.On("Enabled").Return(true)

	authorities := []string{
		delaymsgtypes.ModuleAddress.String(),
		lib.GovModuleAddress.String(),
	}
	k := keeper.NewKeeper(
		cdc,
		epochsKeeper,
		storeKey,
		transientStoreKey,
		authorities,
	)

	return k, storeKey
}
