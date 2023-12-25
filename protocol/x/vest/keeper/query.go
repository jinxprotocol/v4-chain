package keeper

import (
	"github.com/jinxprotocol/v4-chain/protocol/x/vest/types"
)

var _ types.QueryServer = Keeper{}
