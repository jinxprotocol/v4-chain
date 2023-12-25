package keeper

import (
	"github.com/jinxprotocol/v4-chain/protocol/x/assets/types"
)

var _ types.QueryServer = Keeper{}
