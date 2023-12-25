package msgs

import (
	"github.com/jinxprotocol/v4-chain/protocol/lib"
)

var (
	// AllTypeMessages is a list of all messages and types that are used in the app.
	// This list comes from the app's `InterfaceRegistry`.
	AllTypeMessages = map[string]struct{}{
		// auth
		"/cosmos.auth.v1beta1.BaseAccount":      {},
		"/cosmos.auth.v1beta1.ModuleAccount":    {},
		"/cosmos.auth.v1beta1.ModuleCredential": {},
		"/cosmos.auth.v1beta1.MsgUpdateParams":  {},

		// bank
		"/cosmos.bank.v1beta1.MsgMultiSend":              {},
		"/cosmos.bank.v1beta1.MsgMultiSendResponse":      {},
		"/cosmos.bank.v1beta1.MsgSend":                   {},
		"/cosmos.bank.v1beta1.MsgSendResponse":           {},
		"/cosmos.bank.v1beta1.MsgSetSendEnabled":         {},
		"/cosmos.bank.v1beta1.MsgSetSendEnabledResponse": {},
		"/cosmos.bank.v1beta1.MsgUpdateParams":           {},
		"/cosmos.bank.v1beta1.MsgUpdateParamsResponse":   {},
		"/cosmos.bank.v1beta1.SendAuthorization":         {},
		"/cosmos.bank.v1beta1.Supply":                    {},

		// consensus
		"/cosmos.consensus.v1.MsgUpdateParams":         {},
		"/cosmos.consensus.v1.MsgUpdateParamsResponse": {},

		// crisis
		"/cosmos.crisis.v1beta1.MsgUpdateParams":            {},
		"/cosmos.crisis.v1beta1.MsgUpdateParamsResponse":    {},
		"/cosmos.crisis.v1beta1.MsgVerifyInvariant":         {},
		"/cosmos.crisis.v1beta1.MsgVerifyInvariantResponse": {},

		// crypto
		"/cosmos.crypto.ed25519.PrivKey":            {},
		"/cosmos.crypto.ed25519.PubKey":             {},
		"/cosmos.crypto.multisig.LegacyAminoPubKey": {},
		"/cosmos.crypto.secp256k1.PrivKey":          {},
		"/cosmos.crypto.secp256k1.PubKey":           {},
		"/cosmos.crypto.secp256r1.PubKey":           {},

		// distribution
		"/cosmos.distribution.v1beta1.CommunityPoolSpendProposal":             {},
		"/cosmos.distribution.v1beta1.MsgCommunityPoolSpend":                  {},
		"/cosmos.distribution.v1beta1.MsgCommunityPoolSpendResponse":          {},
		"/cosmos.distribution.v1beta1.MsgFundCommunityPool":                   {},
		"/cosmos.distribution.v1beta1.MsgFundCommunityPoolResponse":           {},
		"/cosmos.distribution.v1beta1.MsgSetWithdrawAddress":                  {},
		"/cosmos.distribution.v1beta1.MsgSetWithdrawAddressResponse":          {},
		"/cosmos.distribution.v1beta1.MsgUpdateParams":                        {},
		"/cosmos.distribution.v1beta1.MsgUpdateParamsResponse":                {},
		"/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward":             {},
		"/cosmos.distribution.v1beta1.MsgWithdrawDelegatorRewardResponse":     {},
		"/cosmos.distribution.v1beta1.MsgWithdrawValidatorCommission":         {},
		"/cosmos.distribution.v1beta1.MsgWithdrawValidatorCommissionResponse": {},

		// evidence
		"/cosmos.evidence.v1beta1.Equivocation":              {},
		"/cosmos.evidence.v1beta1.MsgSubmitEvidence":         {},
		"/cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse": {},

		// feegrant
		"/cosmos.feegrant.v1beta1.AllowedMsgAllowance":        {},
		"/cosmos.feegrant.v1beta1.BasicAllowance":             {},
		"/cosmos.feegrant.v1beta1.MsgGrantAllowance":          {},
		"/cosmos.feegrant.v1beta1.MsgGrantAllowanceResponse":  {},
		"/cosmos.feegrant.v1beta1.MsgRevokeAllowance":         {},
		"/cosmos.feegrant.v1beta1.MsgRevokeAllowanceResponse": {},
		"/cosmos.feegrant.v1beta1.PeriodicAllowance":          {},

		// gov
		"/cosmos.gov.v1.MsgDeposit":                     {},
		"/cosmos.gov.v1.MsgDepositResponse":             {},
		"/cosmos.gov.v1.MsgExecLegacyContent":           {},
		"/cosmos.gov.v1.MsgExecLegacyContentResponse":   {},
		"/cosmos.gov.v1.MsgSubmitProposal":              {},
		"/cosmos.gov.v1.MsgSubmitProposalResponse":      {},
		"/cosmos.gov.v1.MsgUpdateParams":                {},
		"/cosmos.gov.v1.MsgUpdateParamsResponse":        {},
		"/cosmos.gov.v1.MsgVote":                        {},
		"/cosmos.gov.v1.MsgVoteResponse":                {},
		"/cosmos.gov.v1.MsgVoteWeighted":                {},
		"/cosmos.gov.v1.MsgVoteWeightedResponse":        {},
		"/cosmos.gov.v1beta1.MsgDeposit":                {},
		"/cosmos.gov.v1beta1.MsgDepositResponse":        {},
		"/cosmos.gov.v1beta1.MsgSubmitProposal":         {},
		"/cosmos.gov.v1beta1.MsgSubmitProposalResponse": {},
		"/cosmos.gov.v1beta1.MsgVote":                   {},
		"/cosmos.gov.v1beta1.MsgVoteResponse":           {},
		"/cosmos.gov.v1beta1.MsgVoteWeighted":           {},
		"/cosmos.gov.v1beta1.MsgVoteWeightedResponse":   {},
		"/cosmos.gov.v1beta1.TextProposal":              {},

		// params
		"/cosmos.params.v1beta1.ParameterChangeProposal": {},

		// slashing
		"/cosmos.slashing.v1beta1.MsgUnjail":               {},
		"/cosmos.slashing.v1beta1.MsgUnjailResponse":       {},
		"/cosmos.slashing.v1beta1.MsgUpdateParams":         {},
		"/cosmos.slashing.v1beta1.MsgUpdateParamsResponse": {},

		// staking
		"/cosmos.staking.v1beta1.MsgBeginRedelegate":                   {},
		"/cosmos.staking.v1beta1.MsgBeginRedelegateResponse":           {},
		"/cosmos.staking.v1beta1.MsgCancelUnbondingDelegation":         {},
		"/cosmos.staking.v1beta1.MsgCancelUnbondingDelegationResponse": {},
		"/cosmos.staking.v1beta1.MsgCreateValidator":                   {},
		"/cosmos.staking.v1beta1.MsgCreateValidatorResponse":           {},
		"/cosmos.staking.v1beta1.MsgDelegate":                          {},
		"/cosmos.staking.v1beta1.MsgDelegateResponse":                  {},
		"/cosmos.staking.v1beta1.MsgEditValidator":                     {},
		"/cosmos.staking.v1beta1.MsgEditValidatorResponse":             {},
		"/cosmos.staking.v1beta1.MsgUndelegate":                        {},
		"/cosmos.staking.v1beta1.MsgUndelegateResponse":                {},
		"/cosmos.staking.v1beta1.MsgUpdateParams":                      {},
		"/cosmos.staking.v1beta1.MsgUpdateParamsResponse":              {},
		"/cosmos.staking.v1beta1.StakeAuthorization":                   {},

		// tx
		"/cosmos.tx.v1beta1.Tx": {},

		// upgrade
		"/cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal": {},
		"/cosmos.upgrade.v1beta1.MsgCancelUpgrade":              {},
		"/cosmos.upgrade.v1beta1.MsgCancelUpgradeResponse":      {},
		"/cosmos.upgrade.v1beta1.MsgSoftwareUpgrade":            {},
		"/cosmos.upgrade.v1beta1.MsgSoftwareUpgradeResponse":    {},
		"/cosmos.upgrade.v1beta1.SoftwareUpgradeProposal":       {},

		// blocktime
		"/jinxprotocol.blocktime.MsgUpdateDowntimeParams":         {},
		"/jinxprotocol.blocktime.MsgUpdateDowntimeParamsResponse": {},

		// bridge
		"/jinxprotocol.bridge.MsgAcknowledgeBridges":          {},
		"/jinxprotocol.bridge.MsgAcknowledgeBridgesResponse":  {},
		"/jinxprotocol.bridge.MsgCompleteBridge":              {},
		"/jinxprotocol.bridge.MsgCompleteBridgeResponse":      {},
		"/jinxprotocol.bridge.MsgUpdateEventParams":           {},
		"/jinxprotocol.bridge.MsgUpdateEventParamsResponse":   {},
		"/jinxprotocol.bridge.MsgUpdateProposeParams":         {},
		"/jinxprotocol.bridge.MsgUpdateProposeParamsResponse": {},
		"/jinxprotocol.bridge.MsgUpdateSafetyParams":          {},
		"/jinxprotocol.bridge.MsgUpdateSafetyParamsResponse":  {},

		// clob
		"/jinxprotocol.clob.MsgCancelOrder":                                {},
		"/jinxprotocol.clob.MsgCancelOrderResponse":                        {},
		"/jinxprotocol.clob.MsgCreateClobPair":                             {},
		"/jinxprotocol.clob.MsgCreateClobPairResponse":                     {},
		"/jinxprotocol.clob.MsgPlaceOrder":                                 {},
		"/jinxprotocol.clob.MsgPlaceOrderResponse":                         {},
		"/jinxprotocol.clob.MsgProposedOperations":                         {},
		"/jinxprotocol.clob.MsgProposedOperationsResponse":                 {},
		"/jinxprotocol.clob.MsgUpdateBlockRateLimitConfiguration":          {},
		"/jinxprotocol.clob.MsgUpdateBlockRateLimitConfigurationResponse":  {},
		"/jinxprotocol.clob.MsgUpdateClobPair":                             {},
		"/jinxprotocol.clob.MsgUpdateClobPairResponse":                     {},
		"/jinxprotocol.clob.MsgUpdateEquityTierLimitConfiguration":         {},
		"/jinxprotocol.clob.MsgUpdateEquityTierLimitConfigurationResponse": {},
		"/jinxprotocol.clob.MsgUpdateLiquidationsConfig":                   {},
		"/jinxprotocol.clob.MsgUpdateLiquidationsConfigResponse":           {},

		// delaymsg
		"/jinxprotocol.delaymsg.MsgDelayMessage":         {},
		"/jinxprotocol.delaymsg.MsgDelayMessageResponse": {},

		// feetiers
		"/jinxprotocol.feetiers.MsgUpdatePerpetualFeeParams":         {},
		"/jinxprotocol.feetiers.MsgUpdatePerpetualFeeParamsResponse": {},

		// perpetuals
		"/jinxprotocol.perpetuals.MsgAddPremiumVotes":               {},
		"/jinxprotocol.perpetuals.MsgAddPremiumVotesResponse":       {},
		"/jinxprotocol.perpetuals.MsgCreatePerpetual":               {},
		"/jinxprotocol.perpetuals.MsgCreatePerpetualResponse":       {},
		"/jinxprotocol.perpetuals.MsgSetLiquidityTier":              {},
		"/jinxprotocol.perpetuals.MsgSetLiquidityTierResponse":      {},
		"/jinxprotocol.perpetuals.MsgUpdateParams":                  {},
		"/jinxprotocol.perpetuals.MsgUpdateParamsResponse":          {},
		"/jinxprotocol.perpetuals.MsgUpdatePerpetualParams":         {},
		"/jinxprotocol.perpetuals.MsgUpdatePerpetualParamsResponse": {},

		// prices
		"/jinxprotocol.prices.MsgCreateOracleMarket":         {},
		"/jinxprotocol.prices.MsgCreateOracleMarketResponse": {},
		"/jinxprotocol.prices.MsgUpdateMarketPrices":         {},
		"/jinxprotocol.prices.MsgUpdateMarketPricesResponse": {},
		"/jinxprotocol.prices.MsgUpdateMarketParam":          {},
		"/jinxprotocol.prices.MsgUpdateMarketParamResponse":  {},

		// sending
		"/jinxprotocol.sending.MsgCreateTransfer":                  {},
		"/jinxprotocol.sending.MsgCreateTransferResponse":          {},
		"/jinxprotocol.sending.MsgDepositToSubaccount":             {},
		"/jinxprotocol.sending.MsgDepositToSubaccountResponse":     {},
		"/jinxprotocol.sending.MsgWithdrawFromSubaccount":          {},
		"/jinxprotocol.sending.MsgWithdrawFromSubaccountResponse":  {},
		"/jinxprotocol.sending.MsgSendFromModuleToAccount":         {},
		"/jinxprotocol.sending.MsgSendFromModuleToAccountResponse": {},

		// stats
		"/jinxprotocol.stats.MsgUpdateParams":         {},
		"/jinxprotocol.stats.MsgUpdateParamsResponse": {},

		// vest
		"/jinxprotocol.vest.MsgSetVestEntry":            {},
		"/jinxprotocol.vest.MsgSetVestEntryResponse":    {},
		"/jinxprotocol.vest.MsgDeleteVestEntry":         {},
		"/jinxprotocol.vest.MsgDeleteVestEntryResponse": {},

		// rewards
		"/jinxprotocol.rewards.MsgUpdateParams":         {},
		"/jinxprotocol.rewards.MsgUpdateParamsResponse": {},

		// ibc.applications
		"/ibc.applications.transfer.v1.MsgTransfer":           {},
		"/ibc.applications.transfer.v1.MsgTransferResponse":   {},
		"/ibc.applications.transfer.v1.TransferAuthorization": {},

		// ibc.core.channel
		"/ibc.core.channel.v1.Channel":                        {},
		"/ibc.core.channel.v1.Counterparty":                   {},
		"/ibc.core.channel.v1.MsgAcknowledgement":             {},
		"/ibc.core.channel.v1.MsgAcknowledgementResponse":     {},
		"/ibc.core.channel.v1.MsgChannelCloseConfirm":         {},
		"/ibc.core.channel.v1.MsgChannelCloseConfirmResponse": {},
		"/ibc.core.channel.v1.MsgChannelCloseInit":            {},
		"/ibc.core.channel.v1.MsgChannelCloseInitResponse":    {},
		"/ibc.core.channel.v1.MsgChannelOpenAck":              {},
		"/ibc.core.channel.v1.MsgChannelOpenAckResponse":      {},
		"/ibc.core.channel.v1.MsgChannelOpenConfirm":          {},
		"/ibc.core.channel.v1.MsgChannelOpenConfirmResponse":  {},
		"/ibc.core.channel.v1.MsgChannelOpenInit":             {},
		"/ibc.core.channel.v1.MsgChannelOpenInitResponse":     {},
		"/ibc.core.channel.v1.MsgChannelOpenTry":              {},
		"/ibc.core.channel.v1.MsgChannelOpenTryResponse":      {},
		"/ibc.core.channel.v1.MsgRecvPacket":                  {},
		"/ibc.core.channel.v1.MsgRecvPacketResponse":          {},
		"/ibc.core.channel.v1.MsgTimeout":                     {},
		"/ibc.core.channel.v1.MsgTimeoutOnClose":              {},
		"/ibc.core.channel.v1.MsgTimeoutOnCloseResponse":      {},
		"/ibc.core.channel.v1.MsgTimeoutResponse":             {},
		"/ibc.core.channel.v1.Packet":                         {},

		// ibc.core.client
		"/ibc.core.client.v1.ClientUpdateProposal":          {},
		"/ibc.core.client.v1.Height":                        {},
		"/ibc.core.client.v1.MsgCreateClient":               {},
		"/ibc.core.client.v1.MsgCreateClientResponse":       {},
		"/ibc.core.client.v1.MsgSubmitMisbehaviour":         {},
		"/ibc.core.client.v1.MsgSubmitMisbehaviourResponse": {},
		"/ibc.core.client.v1.MsgUpdateClient":               {},
		"/ibc.core.client.v1.MsgUpdateClientResponse":       {},
		"/ibc.core.client.v1.MsgUpgradeClient":              {},
		"/ibc.core.client.v1.MsgUpgradeClientResponse":      {},
		"/ibc.core.client.v1.UpgradeProposal":               {},

		// ibc.core.commitment
		"/ibc.core.commitment.v1.MerklePath":   {},
		"/ibc.core.commitment.v1.MerklePrefix": {},
		"/ibc.core.commitment.v1.MerkleProof":  {},
		"/ibc.core.commitment.v1.MerkleRoot":   {},

		// ibc.core.connection
		"/ibc.core.connection.v1.ConnectionEnd":                    {},
		"/ibc.core.connection.v1.Counterparty":                     {},
		"/ibc.core.connection.v1.MsgConnectionOpenAck":             {},
		"/ibc.core.connection.v1.MsgConnectionOpenAckResponse":     {},
		"/ibc.core.connection.v1.MsgConnectionOpenConfirm":         {},
		"/ibc.core.connection.v1.MsgConnectionOpenConfirmResponse": {},
		"/ibc.core.connection.v1.MsgConnectionOpenInit":            {},
		"/ibc.core.connection.v1.MsgConnectionOpenInitResponse":    {},
		"/ibc.core.connection.v1.MsgConnectionOpenTry":             {},
		"/ibc.core.connection.v1.MsgConnectionOpenTryResponse":     {},
		"/ibc.core.connection.v1.Version":                          {},

		// ibc.lightclients
		"/ibc.lightclients.localhost.v2.ClientState":     {},
		"/ibc.lightclients.tendermint.v1.ClientState":    {},
		"/ibc.lightclients.tendermint.v1.ConsensusState": {},
		"/ibc.lightclients.tendermint.v1.Header":         {},
		"/ibc.lightclients.tendermint.v1.Misbehaviour":   {},

		// ica messages
		// Note: the `interchain_accounts.controller` messages are not actually used by the app,
		// since ICA Controller Keeper is initialized as nil.
		// However, since the ica.AppModuleBasic{} needs to be passed to basic_mananger as a whole, these messages
		// registered in the interface registry.
		"/ibc.applications.interchain_accounts.v1.InterchainAccount":                       {},
		"/ibc.applications.interchain_accounts.controller.v1.MsgSendTx":                    {},
		"/ibc.applications.interchain_accounts.controller.v1.MsgRegisterInterchainAccount": {},
	}

	// DisallowMsgs are messages that cannot be externally submitted.
	DisallowMsgs = lib.MergeAllMapsMustHaveDistinctKeys(
		AppInjectedMsgSamples,
		InternalMsgSamplesAll,
		NestedMsgSamples,
		UnsupportedMsgSamples,
	)

	// AllowMsgs are messages that can be externally submitted.
	AllowMsgs = NormalMsgs
)
