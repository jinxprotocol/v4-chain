package keeper

import (
	"github.com/jinxprotocol/v4-chain/protocol/x/delaymsg/types"
)

var _ types.QueryServer = Keeper{}
