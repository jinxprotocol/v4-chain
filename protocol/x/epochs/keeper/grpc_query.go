package keeper

import (
	"github.com/jinxprotocol/v4-chain/protocol/x/epochs/types"
)

var _ types.QueryServer = Keeper{}
