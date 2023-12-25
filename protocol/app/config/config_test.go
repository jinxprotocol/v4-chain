package config_test

import (
	"testing"

	"github.com/jinxprotocol/v4-chain/protocol/app/config"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestSetupConfig_SealsConfig(t *testing.T) {
	sdkConfig := sdk.GetConfig()

	// A successful set confirms the config is not yet sealed
	sdkConfig.SetPurpose(0)
	require.Equal(t, uint32(0), sdkConfig.GetPurpose(), "Expected purpose to match set value")

	// Should set default app values and seal the config
	config.SetupConfig()

	require.Panicsf(t, func() { sdkConfig.SetPurpose(0) }, "Expected config to be sealed after SetupConfig")
}

func TestSetAddressPrefixes(t *testing.T) {
	sdkConfig := sdk.GetConfig()

	require.Equal(t, "jinx", sdkConfig.GetBech32AccountAddrPrefix())
	require.Equal(t, "jinxpub", sdkConfig.GetBech32AccountPubPrefix())

	require.Equal(t, "jinxvaloper", sdkConfig.GetBech32ValidatorAddrPrefix())
	require.Equal(t, "jinxvaloperpub", sdkConfig.GetBech32ValidatorPubPrefix())

	require.Equal(t, "jinxvalcons", sdkConfig.GetBech32ConsensusAddrPrefix())
	require.Equal(t, "jinxvalconspub", sdkConfig.GetBech32ConsensusPubPrefix())
}
