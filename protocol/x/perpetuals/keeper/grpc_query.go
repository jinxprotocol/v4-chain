package keeper

import (
	"github.com/jinxprotocol/v4-chain/protocol/x/perpetuals/types"
)

var _ types.QueryServer = Keeper{}
