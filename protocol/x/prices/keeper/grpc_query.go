package keeper

import (
	"github.com/jinxprotocol/v4-chain/protocol/x/prices/types"
)

var _ types.QueryServer = Keeper{}
