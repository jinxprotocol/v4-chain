package keeper

import (
	"testing"

	indexer_manager "github.com/jinxprotocol/v4-chain/protocol/indexer/indexer_manager"

	tmdb "github.com/cometbft/cometbft-db"
	sdktest "github.com/jinxprotocol/v4-chain/protocol/testutil/sdk"
	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type GenesisInitializer interface {
	InitializeForGenesis(ctx sdk.Context)
}

type callback func(
	db *tmdb.MemDB,
	registry codectypes.InterfaceRegistry,
	cdc *codec.ProtoCodec,
	stateStore storetypes.CommitMultiStore,
	transientStoreKey storetypes.StoreKey,
) []GenesisInitializer

func initKeepers(t testing.TB, cb callback) sdk.Context {
	ctx, stateStore, db := sdktest.NewSdkContextWithMultistore()
	// Mount transient store for indexer events, shared by all keepers that emit indexer events.
	transientStoreKey := sdk.NewTransientStoreKey(indexer_manager.IndexerEventsKey)
	stateStore.MountStoreWithDB(transientStoreKey, storetypes.StoreTypeTransient, db)
	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	initializers := cb(db, registry, cdc, stateStore, transientStoreKey)

	require.NoError(t, stateStore.LoadLatestVersion())

	for _, i := range initializers {
		i.InitializeForGenesis(ctx)
	}

	return ctx
}
