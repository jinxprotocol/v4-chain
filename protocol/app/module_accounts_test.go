package app_test

import (
	"testing"

	icatypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	"github.com/jinxprotocol/v4-chain/protocol/app"
	bridgemoduletypes "github.com/jinxprotocol/v4-chain/protocol/x/bridge/types"
	clobmoduletypes "github.com/jinxprotocol/v4-chain/protocol/x/clob/types"
	delaymsgtypes "github.com/jinxprotocol/v4-chain/protocol/x/delaymsg/types"
	rewardsmoduletypes "github.com/jinxprotocol/v4-chain/protocol/x/rewards/types"
	satypes "github.com/jinxprotocol/v4-chain/protocol/x/subaccounts/types"
	vestmoduletypes "github.com/jinxprotocol/v4-chain/protocol/x/vest/types"
	"github.com/stretchr/testify/require"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func TestModuleAccountsToAddresses(t *testing.T) {
	expectedModuleAccToAddresses := map[string]string{
		authtypes.FeeCollectorName:                   "jinx17xpfvakm2amg962yls6f84z3kell8c5leqdyt2",
		bridgemoduletypes.ModuleName:                 "jinx1zlefkpe3g0vvm9a4h0jf9000lmqutlh94a45jv",
		distrtypes.ModuleName:                        "jinx1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8wx2cfg",
		stakingtypes.BondedPoolName:                  "jinx1fl48vsnmsdzcv85q5d2q4z5ajdha8yu3uz8teq",
		stakingtypes.NotBondedPoolName:               "jinx1tygms3xhhs3yv487phx3dw4a95jn7t7lgzm605",
		govtypes.ModuleName:                          "jinx10d07y265gmmuvt4z0w9aw880jnsr700j5g6x5y",
		ibctransfertypes.ModuleName:                  "jinx1yl6hdjhmkf37639730gffanpzndzdpmh8xcdh5",
		satypes.ModuleName:                           "jinx1v88c3xv9xyv3eetdx0tvcmq7ung3dywp5upwc6",
		clobmoduletypes.InsuranceFundName:            "jinx1c7ptc87hkd54e3r7zjy92q29xkq7t79waxhcpq",
		rewardsmoduletypes.TreasuryAccountName:       "jinx16wrau2x4tsg033xfrrdpae6kxfn9kyueysn4sp",
		rewardsmoduletypes.VesterAccountName:         "jinx1ltyc6y4skclzafvpznpt2qjwmfwgsndp458rmp",
		vestmoduletypes.CommunityTreasuryAccountName: "jinx15ztc7xy42tn2ukkc0qjthkucw9ac63pgp70urn",
		vestmoduletypes.CommunityVesterAccountName:   "jinx1wxje320an3karyc6mjw4zghs300dmrjkwn7xtk",
		delaymsgtypes.ModuleName:                     "jinx1mkkvp26dngu6n8rmalaxyp3gwkjuzztqn3pafr",
		icatypes.ModuleName:                          "jinx1vlthgax23ca9syk7xgaz347xmf4nunefw3cnv8",
	}

	require.True(t, len(expectedModuleAccToAddresses) == len(app.GetMaccPerms()))
	for acc, address := range expectedModuleAccToAddresses {
		expectedAddr := authtypes.NewModuleAddress(acc).String()
		require.Equal(t, address, expectedAddr, "module (%v) should have address (%s)", acc, expectedAddr)
	}
}

func TestBlockedAddresses(t *testing.T) {
	expectedBlockedAddresses := map[string]bool{
		"jinx17xpfvakm2amg962yls6f84z3kell8c5leqdyt2": true,
		"jinx1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8wx2cfg": true,
		"jinx1tygms3xhhs3yv487phx3dw4a95jn7t7lgzm605": true,
		"jinx1fl48vsnmsdzcv85q5d2q4z5ajdha8yu3uz8teq": true,
		"jinx1yl6hdjhmkf37639730gffanpzndzdpmh8xcdh5": true,
		"jinx1vlthgax23ca9syk7xgaz347xmf4nunefw3cnv8": true,
	}
	require.Equal(t, expectedBlockedAddresses, app.BlockedAddresses())
}

func TestMaccPerms(t *testing.T) {
	maccPerms := app.GetMaccPerms()
	expectedMaccPerms := map[string][]string{
		"bonded_tokens_pool":     {"burner", "staking"},
		"bridge":                 {"minter"},
		"distribution":           nil,
		"fee_collector":          nil,
		"gov":                    {"burner"},
		"insurance_fund":         nil,
		"not_bonded_tokens_pool": {"burner", "staking"},
		"subaccounts":            nil,
		"transfer":               {"minter", "burner"},
		"interchainaccounts":     nil,
		"rewards_treasury":       nil,
		"rewards_vester":         nil,
		"community_treasury":     nil,
		"community_vester":       nil,
		"delaymsg":               nil,
	}
	require.Equal(t, expectedMaccPerms, maccPerms, "default macc perms list does not match expected")
}

func TestModuleAccountAddrs(t *testing.T) {
	expectedModuleAccAddresses := map[string]bool{
		"jinx17xpfvakm2amg962yls6f84z3kell8c5leqdyt2": true, // x/auth.FeeCollector
		"jinx1zlefkpe3g0vvm9a4h0jf9000lmqutlh94a45jv": true, // x/bridge
		"jinx1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8wx2cfg": true, // x/distribution
		"jinx1fl48vsnmsdzcv85q5d2q4z5ajdha8yu3uz8teq": true, // x/staking.bondedPool
		"jinx1tygms3xhhs3yv487phx3dw4a95jn7t7lgzm605": true, // x/staking.notBondedPool
		"jinx10d07y265gmmuvt4z0w9aw880jnsr700j5g6x5y": true, // x/ gov
		"jinx1yl6hdjhmkf37639730gffanpzndzdpmh8xcdh5": true, // ibc transfer
		"jinx1vlthgax23ca9syk7xgaz347xmf4nunefw3cnv8": true, // interchainaccounts
		"jinx1v88c3xv9xyv3eetdx0tvcmq7ung3dywp5upwc6": true, // x/subaccount
		"jinx1c7ptc87hkd54e3r7zjy92q29xkq7t79waxhcpq": true, // x/clob.insuranceFund
		"jinx16wrau2x4tsg033xfrrdpae6kxfn9kyueysn4sp": true, // x/rewards.treasury
		"jinx1ltyc6y4skclzafvpznpt2qjwmfwgsndp458rmp": true, // x/rewards.vester
		"jinx15ztc7xy42tn2ukkc0qjthkucw9ac63pgp70urn": true, // x/vest.communityTreasury
		"jinx1wxje320an3karyc6mjw4zghs300dmrjkwn7xtk": true, // x/vest.communityVester
		"jinx1mkkvp26dngu6n8rmalaxyp3gwkjuzztqn3pafr": true, // x/delaymsg
	}

	require.Equal(t, expectedModuleAccAddresses, app.ModuleAccountAddrs())
}
