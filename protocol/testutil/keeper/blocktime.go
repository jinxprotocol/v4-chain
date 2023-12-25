package keeper

import (
	tmdb "github.com/cometbft/cometbft-db"
	"github.com/jinxprotocol/v4-chain/protocol/lib"
	"github.com/jinxprotocol/v4-chain/protocol/x/blocktime/keeper"
	"github.com/jinxprotocol/v4-chain/protocol/x/blocktime/types"
	delaymsgtypes "github.com/jinxprotocol/v4-chain/protocol/x/delaymsg/types"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func createBlockTimeKeeper(
	stateStore storetypes.CommitMultiStore,
	db *tmdb.MemDB,
	cdc *codec.ProtoCodec,
) (*keeper.Keeper, storetypes.StoreKey) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)

	authorities := []string{
		delaymsgtypes.ModuleAddress.String(),
		lib.GovModuleAddress.String(),
	}
	k := keeper.NewKeeper(
		cdc,
		storeKey,
		authorities,
	)

	return k, storeKey
}
