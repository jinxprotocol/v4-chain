package v_3_0_0

import (
	icahosttypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/host/types"
	"github.com/jinxprotocol/v4-chain/protocol/app/upgrades"
	ratelimittypes "github.com/jinxprotocol/v4-chain/protocol/x/ratelimit/types"

	store "github.com/cosmos/cosmos-sdk/store/types"
)

const (
	UpgradeName = "v3.0.0"
)

var Upgrade = upgrades.Upgrade{
	UpgradeName: UpgradeName,
	StoreUpgrades: store.StoreUpgrades{
		Added: []string{
			ratelimittypes.StoreKey,
			icahosttypes.StoreKey,
		},
	},
}
