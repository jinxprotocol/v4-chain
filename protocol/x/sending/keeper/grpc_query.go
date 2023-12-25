package keeper

import (
	"github.com/jinxprotocol/v4-chain/protocol/x/sending/types"
)

var _ types.QueryServer = Keeper{}
