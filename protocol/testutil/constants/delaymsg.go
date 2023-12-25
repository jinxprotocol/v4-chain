package constants

import (
	bridgetypes "github.com/jinxprotocol/v4-chain/protocol/x/bridge/types"
	"github.com/jinxprotocol/v4-chain/protocol/x/delaymsg/types"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	// MsgCompleteBridge is an example of an expected Msg type in the delaymsg module.
	TestMsg1 = &bridgetypes.MsgCompleteBridge{
		Authority: types.ModuleAddress.String(),
		Event: bridgetypes.BridgeEvent{
			Id: 1,
		},
	}
	TestMsg2 = &bridgetypes.MsgCompleteBridge{
		Authority: types.ModuleAddress.String(),
		Event: bridgetypes.BridgeEvent{
			Id: 2,
		},
	}
	TestMsg3 = &bridgetypes.MsgCompleteBridge{
		Authority: types.ModuleAddress.String(),
		Event: bridgetypes.BridgeEvent{
			Id: 3,
		},
	}
	NoHandlerMsg = &testdata.TestMsg{Signers: []string{"invalid - no module handles this message"}}

	AllMsgs = []sdk.Msg{TestMsg1, TestMsg2, TestMsg3}
)
