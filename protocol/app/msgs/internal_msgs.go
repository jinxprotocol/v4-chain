package msgs

import (
	"github.com/jinxprotocol/v4-chain/protocol/lib"
	blocktime "github.com/jinxprotocol/v4-chain/protocol/x/blocktime/types"
	bridge "github.com/jinxprotocol/v4-chain/protocol/x/bridge/types"
	clob "github.com/jinxprotocol/v4-chain/protocol/x/clob/types"
	delaymsg "github.com/jinxprotocol/v4-chain/protocol/x/delaymsg/types"
	feetiers "github.com/jinxprotocol/v4-chain/protocol/x/feetiers/types"
	perpetuals "github.com/jinxprotocol/v4-chain/protocol/x/perpetuals/types"
	prices "github.com/jinxprotocol/v4-chain/protocol/x/prices/types"
	rewards "github.com/jinxprotocol/v4-chain/protocol/x/rewards/types"
	sending "github.com/jinxprotocol/v4-chain/protocol/x/sending/types"
	stats "github.com/jinxprotocol/v4-chain/protocol/x/stats/types"
	vest "github.com/jinxprotocol/v4-chain/protocol/x/vest/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	auth "github.com/cosmos/cosmos-sdk/x/auth/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	consensus "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisis "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distribution "github.com/cosmos/cosmos-sdk/x/distribution/types"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	slashing "github.com/cosmos/cosmos-sdk/x/slashing/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgrade "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

var (
	// InternalMsgSamplesAll are msgs that are used only used internally.
	InternalMsgSamplesAll = lib.MergeAllMapsMustHaveDistinctKeys(InternalMsgSamplesGovAuth)

	// InternalMsgSamplesGovAuth are msgs that are used only used internally.
	// GovAuth means that these messages must originate from the gov module and
	// signed by gov module account.
	InternalMsgSamplesGovAuth = map[string]sdk.Msg{
		// ------- CosmosSDK default modules
		// auth
		"/cosmos.auth.v1beta1.MsgUpdateParams": &auth.MsgUpdateParams{},

		// bank
		"/cosmos.bank.v1beta1.MsgSetSendEnabled":         &bank.MsgSetSendEnabled{},
		"/cosmos.bank.v1beta1.MsgSetSendEnabledResponse": nil,
		"/cosmos.bank.v1beta1.MsgUpdateParams":           &bank.MsgUpdateParams{},
		"/cosmos.bank.v1beta1.MsgUpdateParamsResponse":   nil,

		// consensus
		"/cosmos.consensus.v1.MsgUpdateParams":         &consensus.MsgUpdateParams{},
		"/cosmos.consensus.v1.MsgUpdateParamsResponse": nil,

		// crisis
		"/cosmos.crisis.v1beta1.MsgUpdateParams":         &crisis.MsgUpdateParams{},
		"/cosmos.crisis.v1beta1.MsgUpdateParamsResponse": nil,

		// distribution
		"/cosmos.distribution.v1beta1.MsgCommunityPoolSpend":         &distribution.MsgCommunityPoolSpend{},
		"/cosmos.distribution.v1beta1.MsgCommunityPoolSpendResponse": nil,
		"/cosmos.distribution.v1beta1.MsgUpdateParams":               &distribution.MsgUpdateParams{},
		"/cosmos.distribution.v1beta1.MsgUpdateParamsResponse":       nil,

		// gov
		"/cosmos.gov.v1.MsgExecLegacyContent":         &gov.MsgExecLegacyContent{},
		"/cosmos.gov.v1.MsgExecLegacyContentResponse": nil,
		"/cosmos.gov.v1.MsgUpdateParams":              &gov.MsgUpdateParams{},
		"/cosmos.gov.v1.MsgUpdateParamsResponse":      nil,

		// slashing
		"/cosmos.slashing.v1beta1.MsgUpdateParams":         &slashing.MsgUpdateParams{},
		"/cosmos.slashing.v1beta1.MsgUpdateParamsResponse": nil,

		// staking
		"/cosmos.staking.v1beta1.MsgUpdateParams":         &staking.MsgUpdateParams{},
		"/cosmos.staking.v1beta1.MsgUpdateParamsResponse": nil,

		// upgrade
		"/cosmos.upgrade.v1beta1.MsgCancelUpgrade":           &upgrade.MsgCancelUpgrade{},
		"/cosmos.upgrade.v1beta1.MsgCancelUpgradeResponse":   nil,
		"/cosmos.upgrade.v1beta1.MsgSoftwareUpgrade":         &upgrade.MsgSoftwareUpgrade{},
		"/cosmos.upgrade.v1beta1.MsgSoftwareUpgradeResponse": nil,

		// ------- Custom modules
		// blocktime
		"/jinxprotocol.blocktime.MsgUpdateDowntimeParams":         &blocktime.MsgUpdateDowntimeParams{},
		"/jinxprotocol.blocktime.MsgUpdateDowntimeParamsResponse": nil,

		// bridge
		"/jinxprotocol.bridge.MsgCompleteBridge":              &bridge.MsgCompleteBridge{},
		"/jinxprotocol.bridge.MsgCompleteBridgeResponse":      nil,
		"/jinxprotocol.bridge.MsgUpdateEventParams":           &bridge.MsgUpdateEventParams{},
		"/jinxprotocol.bridge.MsgUpdateEventParamsResponse":   nil,
		"/jinxprotocol.bridge.MsgUpdateProposeParams":         &bridge.MsgUpdateProposeParams{},
		"/jinxprotocol.bridge.MsgUpdateProposeParamsResponse": nil,
		"/jinxprotocol.bridge.MsgUpdateSafetyParams":          &bridge.MsgUpdateSafetyParams{},
		"/jinxprotocol.bridge.MsgUpdateSafetyParamsResponse":  nil,

		// clob
		"/jinxprotocol.clob.MsgCreateClobPair":                             &clob.MsgCreateClobPair{},
		"/jinxprotocol.clob.MsgCreateClobPairResponse":                     nil,
		"/jinxprotocol.clob.MsgUpdateBlockRateLimitConfiguration":          &clob.MsgUpdateBlockRateLimitConfiguration{},
		"/jinxprotocol.clob.MsgUpdateBlockRateLimitConfigurationResponse":  nil,
		"/jinxprotocol.clob.MsgUpdateClobPair":                             &clob.MsgUpdateClobPair{},
		"/jinxprotocol.clob.MsgUpdateClobPairResponse":                     nil,
		"/jinxprotocol.clob.MsgUpdateEquityTierLimitConfiguration":         &clob.MsgUpdateEquityTierLimitConfiguration{},
		"/jinxprotocol.clob.MsgUpdateEquityTierLimitConfigurationResponse": nil,
		"/jinxprotocol.clob.MsgUpdateLiquidationsConfig":                   &clob.MsgUpdateLiquidationsConfig{},
		"/jinxprotocol.clob.MsgUpdateLiquidationsConfigResponse":           nil,

		// delaymsg
		"/jinxprotocol.delaymsg.MsgDelayMessage":         &delaymsg.MsgDelayMessage{},
		"/jinxprotocol.delaymsg.MsgDelayMessageResponse": nil,

		// feetiers
		"/jinxprotocol.feetiers.MsgUpdatePerpetualFeeParams":         &feetiers.MsgUpdatePerpetualFeeParams{},
		"/jinxprotocol.feetiers.MsgUpdatePerpetualFeeParamsResponse": nil,

		// perpetuals
		"/jinxprotocol.perpetuals.MsgCreatePerpetual":               &perpetuals.MsgCreatePerpetual{},
		"/jinxprotocol.perpetuals.MsgCreatePerpetualResponse":       nil,
		"/jinxprotocol.perpetuals.MsgSetLiquidityTier":              &perpetuals.MsgSetLiquidityTier{},
		"/jinxprotocol.perpetuals.MsgSetLiquidityTierResponse":      nil,
		"/jinxprotocol.perpetuals.MsgUpdateParams":                  &perpetuals.MsgUpdateParams{},
		"/jinxprotocol.perpetuals.MsgUpdateParamsResponse":          nil,
		"/jinxprotocol.perpetuals.MsgUpdatePerpetualParams":         &perpetuals.MsgUpdatePerpetualParams{},
		"/jinxprotocol.perpetuals.MsgUpdatePerpetualParamsResponse": nil,

		// prices
		"/jinxprotocol.prices.MsgCreateOracleMarket":         &prices.MsgCreateOracleMarket{},
		"/jinxprotocol.prices.MsgCreateOracleMarketResponse": nil,
		"/jinxprotocol.prices.MsgUpdateMarketParam":          &prices.MsgUpdateMarketParam{},
		"/jinxprotocol.prices.MsgUpdateMarketParamResponse":  nil,

		// rewards
		"/jinxprotocol.rewards.MsgUpdateParams":         &rewards.MsgUpdateParams{},
		"/jinxprotocol.rewards.MsgUpdateParamsResponse": nil,

		// sending
		"/jinxprotocol.sending.MsgSendFromModuleToAccount":         &sending.MsgSendFromModuleToAccount{},
		"/jinxprotocol.sending.MsgSendFromModuleToAccountResponse": nil,

		// stats
		"/jinxprotocol.stats.MsgUpdateParams":         &stats.MsgUpdateParams{},
		"/jinxprotocol.stats.MsgUpdateParamsResponse": nil,

		// vest
		"/jinxprotocol.vest.MsgSetVestEntry":            &vest.MsgSetVestEntry{},
		"/jinxprotocol.vest.MsgSetVestEntryResponse":    nil,
		"/jinxprotocol.vest.MsgDeleteVestEntry":         &vest.MsgDeleteVestEntry{},
		"/jinxprotocol.vest.MsgDeleteVestEntryResponse": nil,
	}
)
