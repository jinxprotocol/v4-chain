package msgs_test

import (
	"sort"
	"testing"

	"github.com/jinxprotocol/v4-chain/protocol/app/msgs"
	"github.com/jinxprotocol/v4-chain/protocol/lib"
	"github.com/stretchr/testify/require"
)

func TestInternalMsgSamples_All_Key(t *testing.T) {
	expectedAllInternalMsgs := lib.MergeAllMapsMustHaveDistinctKeys(msgs.InternalMsgSamplesGovAuth)
	require.Equal(t, expectedAllInternalMsgs, msgs.InternalMsgSamplesAll)
}

func TestInternalMsgSamples_All_Value(t *testing.T) {
	validateMsgValue(t, msgs.InternalMsgSamplesAll)
}

func TestInternalMsgSamples_Gov_Key(t *testing.T) {
	expectedMsgs := []string{
		// auth
		"/cosmos.auth.v1beta1.MsgUpdateParams",

		// bank
		"/cosmos.bank.v1beta1.MsgSetSendEnabled",
		"/cosmos.bank.v1beta1.MsgSetSendEnabledResponse",
		"/cosmos.bank.v1beta1.MsgUpdateParams",
		"/cosmos.bank.v1beta1.MsgUpdateParamsResponse",

		// consensus
		"/cosmos.consensus.v1.MsgUpdateParams",
		"/cosmos.consensus.v1.MsgUpdateParamsResponse",

		// crisis
		"/cosmos.crisis.v1beta1.MsgUpdateParams",
		"/cosmos.crisis.v1beta1.MsgUpdateParamsResponse",

		// distribution
		"/cosmos.distribution.v1beta1.MsgCommunityPoolSpend",
		"/cosmos.distribution.v1beta1.MsgCommunityPoolSpendResponse",
		"/cosmos.distribution.v1beta1.MsgUpdateParams",
		"/cosmos.distribution.v1beta1.MsgUpdateParamsResponse",

		// gov
		"/cosmos.gov.v1.MsgExecLegacyContent",
		"/cosmos.gov.v1.MsgExecLegacyContentResponse",
		"/cosmos.gov.v1.MsgUpdateParams",
		"/cosmos.gov.v1.MsgUpdateParamsResponse",

		// slashing
		"/cosmos.slashing.v1beta1.MsgUpdateParams",
		"/cosmos.slashing.v1beta1.MsgUpdateParamsResponse",

		// staking
		"/cosmos.staking.v1beta1.MsgUpdateParams",
		"/cosmos.staking.v1beta1.MsgUpdateParamsResponse",

		// upgrade
		"/cosmos.upgrade.v1beta1.MsgCancelUpgrade",
		"/cosmos.upgrade.v1beta1.MsgCancelUpgradeResponse",
		"/cosmos.upgrade.v1beta1.MsgSoftwareUpgrade",
		"/cosmos.upgrade.v1beta1.MsgSoftwareUpgradeResponse",

		// blocktime
		"/jinxprotocol.blocktime.MsgUpdateDowntimeParams",
		"/jinxprotocol.blocktime.MsgUpdateDowntimeParamsResponse",

		// bridge
		"/jinxprotocol.bridge.MsgCompleteBridge",
		"/jinxprotocol.bridge.MsgCompleteBridgeResponse",
		"/jinxprotocol.bridge.MsgUpdateEventParams",
		"/jinxprotocol.bridge.MsgUpdateEventParamsResponse",
		"/jinxprotocol.bridge.MsgUpdateProposeParams",
		"/jinxprotocol.bridge.MsgUpdateProposeParamsResponse",
		"/jinxprotocol.bridge.MsgUpdateSafetyParams",
		"/jinxprotocol.bridge.MsgUpdateSafetyParamsResponse",

		// clob
		"/jinxprotocol.clob.MsgCreateClobPair",
		"/jinxprotocol.clob.MsgCreateClobPairResponse",
		"/jinxprotocol.clob.MsgUpdateBlockRateLimitConfiguration",
		"/jinxprotocol.clob.MsgUpdateBlockRateLimitConfigurationResponse",
		"/jinxprotocol.clob.MsgUpdateClobPair",
		"/jinxprotocol.clob.MsgUpdateClobPairResponse",
		"/jinxprotocol.clob.MsgUpdateEquityTierLimitConfiguration",
		"/jinxprotocol.clob.MsgUpdateEquityTierLimitConfigurationResponse",
		"/jinxprotocol.clob.MsgUpdateLiquidationsConfig",
		"/jinxprotocol.clob.MsgUpdateLiquidationsConfigResponse",

		// delaymsg
		"/jinxprotocol.delaymsg.MsgDelayMessage",
		"/jinxprotocol.delaymsg.MsgDelayMessageResponse",

		// feetiers
		"/jinxprotocol.feetiers.MsgUpdatePerpetualFeeParams",
		"/jinxprotocol.feetiers.MsgUpdatePerpetualFeeParamsResponse",

		// perpeutals
		"/jinxprotocol.perpetuals.MsgCreatePerpetual",
		"/jinxprotocol.perpetuals.MsgCreatePerpetualResponse",
		"/jinxprotocol.perpetuals.MsgSetLiquidityTier",
		"/jinxprotocol.perpetuals.MsgSetLiquidityTierResponse",
		"/jinxprotocol.perpetuals.MsgUpdateParams",
		"/jinxprotocol.perpetuals.MsgUpdateParamsResponse",
		"/jinxprotocol.perpetuals.MsgUpdatePerpetualParams",
		"/jinxprotocol.perpetuals.MsgUpdatePerpetualParamsResponse",

		// prices
		"/jinxprotocol.prices.MsgCreateOracleMarket",
		"/jinxprotocol.prices.MsgCreateOracleMarketResponse",
		"/jinxprotocol.prices.MsgUpdateMarketParam",
		"/jinxprotocol.prices.MsgUpdateMarketParamResponse",

		// rewards
		"/jinxprotocol.rewards.MsgUpdateParams",
		"/jinxprotocol.rewards.MsgUpdateParamsResponse",

		// sending
		"/jinxprotocol.sending.MsgSendFromModuleToAccount",
		"/jinxprotocol.sending.MsgSendFromModuleToAccountResponse",

		// stats
		"/jinxprotocol.stats.MsgUpdateParams",
		"/jinxprotocol.stats.MsgUpdateParamsResponse",

		// vest
		"/jinxprotocol.vest.MsgDeleteVestEntry",
		"/jinxprotocol.vest.MsgDeleteVestEntryResponse",
		"/jinxprotocol.vest.MsgSetVestEntry",
		"/jinxprotocol.vest.MsgSetVestEntryResponse",
	}

	require.Equal(t, expectedMsgs, lib.GetSortedKeys[sort.StringSlice](msgs.InternalMsgSamplesGovAuth))
}

func TestInternalMsgSamples_Gov_Value(t *testing.T) {
	validateMsgValue(t, msgs.InternalMsgSamplesGovAuth)
}
