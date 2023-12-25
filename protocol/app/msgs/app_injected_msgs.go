package msgs

import (
	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/jinxprotocol/v4-chain/protocol/testutil/constants"
	bridgetypes "github.com/jinxprotocol/v4-chain/protocol/x/bridge/types"
	clobtypes "github.com/jinxprotocol/v4-chain/protocol/x/clob/types"
	perptypes "github.com/jinxprotocol/v4-chain/protocol/x/perpetuals/types"
	pricestypes "github.com/jinxprotocol/v4-chain/protocol/x/prices/types"
)

var (
	// AppInjectedMsgSamples are msgs that are injected into the block by the proposing validator.
	// These messages are reserved for proposing validator's use only.
	AppInjectedMsgSamples = map[string]sdk.Msg{
		// bridge
		"/jinxprotocol.bridge.MsgAcknowledgeBridges": &bridgetypes.MsgAcknowledgeBridges{
			Events: []bridgetypes.BridgeEvent{
				{
					Id: 0,
					Coin: sdk.NewCoin(
						"bridge-token",
						sdkmath.NewIntFromUint64(1234),
					),
					Address: constants.Alice_Num0.Owner,
				},
			},
		},
		"/jinxprotocol.bridge.MsgAcknowledgeBridgesResponse": nil,

		// clob
		"/jinxprotocol.clob.MsgProposedOperations": &clobtypes.MsgProposedOperations{
			OperationsQueue: make([]clobtypes.OperationRaw, 0),
		},
		"/jinxprotocol.clob.MsgProposedOperationsResponse": nil,

		// perpetuals
		"/jinxprotocol.perpetuals.MsgAddPremiumVotes": &perptypes.MsgAddPremiumVotes{
			Votes: []perptypes.FundingPremium{
				{PerpetualId: 0, PremiumPpm: 1_000},
			},
		},
		"/jinxprotocol.perpetuals.MsgAddPremiumVotesResponse": nil,

		// prices
		"/jinxprotocol.prices.MsgUpdateMarketPrices": &pricestypes.MsgUpdateMarketPrices{
			MarketPriceUpdates: []*pricestypes.MsgUpdateMarketPrices_MarketPrice{
				pricestypes.NewMarketPriceUpdate(constants.MarketId0, 123_000),
			},
		},
		"/jinxprotocol.prices.MsgUpdateMarketPricesResponse": nil,
	}
)
