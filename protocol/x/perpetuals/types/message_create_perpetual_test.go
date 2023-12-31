package types_test

import (
	"testing"

	"github.com/jinxprotocol/v4-chain/protocol/testutil/constants"
	types "github.com/jinxprotocol/v4-chain/protocol/x/perpetuals/types"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	// validAuthority is a valid bech32 address string.
	validAuthority = constants.AliceAccAddress.String()
)

func TestMsgCreatePerpetual_GetSigners(t *testing.T) {
	msg := types.MsgCreatePerpetual{
		Authority: constants.AliceAccAddress.String(),
	}
	require.Equal(t, []sdk.AccAddress{constants.AliceAccAddress}, msg.GetSigners())
}

func TestMsgCreatePerpetual_ValidateBasic(t *testing.T) {
	tests := []struct {
		desc        string
		msg         types.MsgCreatePerpetual
		expectedErr string
	}{
		{
			desc: "Invalid authority",
			msg: types.MsgCreatePerpetual{
				Authority: "",
			},
			expectedErr: "Authority is invalid",
		},
		{
			desc: "Empty ticker",
			msg: types.MsgCreatePerpetual{
				Authority: validAuthority,
				Params: types.PerpetualParams{
					Ticker: "",
				},
			},
			expectedErr: "Ticker must be non-empty string",
		},
		{
			desc: "DefaultFundingPpm >= MaxDefaultFundingPpmAbs",
			msg: types.MsgCreatePerpetual{
				Authority: validAuthority,
				Params: types.PerpetualParams{
					Ticker:            "test",
					DefaultFundingPpm: 100_000_000,
				},
			},
			expectedErr: "DefaultFundingPpm magnitude exceeds maximum value",
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			require.ErrorContains(t, err, tc.expectedErr)
		})
	}
}
