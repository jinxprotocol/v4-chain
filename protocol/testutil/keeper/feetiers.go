package keeper

import (
	"github.com/jinxprotocol/v4-chain/protocol/lib"
	"github.com/jinxprotocol/v4-chain/protocol/mocks"

	tmdb "github.com/cometbft/cometbft-db"
	delaymsgtypes "github.com/jinxprotocol/v4-chain/protocol/x/delaymsg/types"
	"github.com/jinxprotocol/v4-chain/protocol/x/feetiers/keeper"
	"github.com/jinxprotocol/v4-chain/protocol/x/feetiers/types"
	statskeeper "github.com/jinxprotocol/v4-chain/protocol/x/stats/keeper"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func createFeeTiersKeeper(
	stateStore storetypes.CommitMultiStore,
	statsKeeper *statskeeper.Keeper,
	db *tmdb.MemDB,
	cdc *codec.ProtoCodec,
) (*keeper.Keeper, storetypes.StoreKey) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)

	mockMsgSender := &mocks.IndexerMessageSender{}
	mockMsgSender.On("Enabled").Return(true)

	authorities := []string{
		delaymsgtypes.ModuleAddress.String(),
		lib.GovModuleAddress.String(),
	}
	k := keeper.NewKeeper(
		cdc,
		statsKeeper,
		storeKey,
		authorities,
	)

	return k, storeKey
}
