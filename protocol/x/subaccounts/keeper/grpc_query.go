package keeper

import (
	"github.com/jinxprotocol/v4-chain/protocol/x/subaccounts/types"
)

var _ types.QueryServer = Keeper{}
