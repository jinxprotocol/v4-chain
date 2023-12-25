package keeper

import (
	"github.com/jinxprotocol/v4-chain/protocol/x/clob/types"
)

var _ types.QueryServer = Keeper{}
