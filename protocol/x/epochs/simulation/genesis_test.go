package simulation_test

import (
	"encoding/json"
	"testing"

	sdkmath "cosmossdk.io/math"
	testutil_rand "github.com/jinxprotocol/v4-chain/protocol/testutil/rand"
	"github.com/jinxprotocol/v4-chain/protocol/x/epochs/simulation"
	"github.com/jinxprotocol/v4-chain/protocol/x/epochs/types"
	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func TestRandomizedGenState(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)

	r := testutil_rand.NewRand()

	simState := module.SimulationState{
		AppParams:    make(simtypes.AppParams),
		Cdc:          cdc,
		Rand:         r,
		NumBonded:    3,
		Accounts:     simtypes.RandomAccounts(r, 3),
		InitialStake: sdkmath.NewInt(1000),
		GenState:     make(map[string]json.RawMessage),
		GenTimestamp: simtypes.RandTimestamp(r),
	}

	for i := 0; i < 100; i++ {
		simulation.RandomizedGenState(&simState)
		var epochsGenesis types.GenesisState
		simState.Cdc.MustUnmarshalJSON(simState.GenState[types.ModuleName], &epochsGenesis)

		require.True(t, len(epochsGenesis.EpochInfoList) >= 1+3)     // +3 comes from default genesis
		require.True(t, len(epochsGenesis.EpochInfoList) <= 1_000+3) // +3 comes from default genesis

		for _, epochInfo := range epochsGenesis.EpochInfoList {
			require.True(t, len(epochInfo.Name) >= 5)
			require.True(t, len(epochInfo.Name) <= 19)

			require.True(t, epochInfo.Duration >= 1)
			// no need to check duration max value since it's already capped by uint32 type.

			// no need to check `NextTick`, `CurrentEpcoh` or `CurrentEpochStartBlock` value
			// since we randomly generate any uint32 value.

			// no need to check `FastForwardNextTick` value since we generate both true/false.
		}
	}
}
