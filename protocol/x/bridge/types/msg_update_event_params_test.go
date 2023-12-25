package types_test

import (
	"testing"

	"github.com/jinxprotocol/v4-chain/protocol/testutil/constants"
	"github.com/jinxprotocol/v4-chain/protocol/x/bridge/types"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	// validAuthority is a valid bech32 address string.
	validAuthority = constants.AliceAccAddress.String()
)

func TestMsgUpdateEventParams_GetSigners(t *testing.T) {
	msg := types.MsgUpdateEventParams{
		Authority: constants.CarlAccAddress.String(),
	}
	require.Equal(t, []sdk.AccAddress{constants.CarlAccAddress}, msg.GetSigners())
}

func TestMsgUpdateEventParams_ValidateBasic(t *testing.T) {
	tests := map[string]struct {
		msg         types.MsgUpdateEventParams
		expectedErr string
	}{
		"Success": {
			msg: types.MsgUpdateEventParams{
				Authority: validAuthority,
				Params: types.EventParams{
					Denom:      "test-denom",
					EthChainId: 0,
					EthAddress: "test",
				},
			},
		},
		"Failure: invalid Denom param": {
			msg: types.MsgUpdateEventParams{
				Authority: validAuthority,
				Params: types.EventParams{
					Denom:      "2test-denom", // cannot start with number
					EthChainId: 0,
					EthAddress: "test",
				},
			},
			expectedErr: "invalid denom",
		},
		"Failure: invalid EthAddress param": {
			msg: types.MsgUpdateEventParams{
				Authority: validAuthority,
				Params: types.EventParams{
					Denom:      "test-denom",
					EthChainId: 0,
					EthAddress: "", // cannot be empty
				},
			},
			expectedErr: types.ErrInvalidEthAddress.Error(),
		},
		"Failure: empty authority": {
			msg: types.MsgUpdateEventParams{
				Authority: "",
			},
			expectedErr: types.ErrInvalidAuthority.Error(),
		},
		"Failure: invalid authority": {
			msg: types.MsgUpdateEventParams{
				Authority: "jinx1abc",
			},
			expectedErr: types.ErrInvalidAuthority.Error(),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.expectedErr == "" {
				require.NoError(t, err)
			} else {
				require.ErrorContains(t, err, tc.expectedErr)
			}
		})
	}
}
