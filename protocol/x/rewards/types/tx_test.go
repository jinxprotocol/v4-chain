package types_test

import (
	"testing"

	"github.com/jinxprotocol/v4-chain/protocol/testutil/constants"
	"github.com/jinxprotocol/v4-chain/protocol/x/rewards/types"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	validAuthority = constants.BobAccAddress.String()
)

func TestGetSigners(t *testing.T) {
	msg := types.MsgUpdateParams{
		Authority: validAuthority,
	}
	require.Equal(t, []sdk.AccAddress{constants.BobAccAddress}, msg.GetSigners())
}

func TestValidateBasic(t *testing.T) {
	test := map[string]struct {
		msg         types.MsgUpdateParams
		expectedErr error
	}{
		"Success": {
			msg: types.MsgUpdateParams{
				Authority: validAuthority,
				Params: types.Params{
					TreasuryAccount: "treasury_account",
					Denom:           "denom",
				},
			},
		},
		"Failure: Invalid authority": {
			msg: types.MsgUpdateParams{
				Authority: "", // invalid - empty
			},
			expectedErr: types.ErrInvalidAuthority,
		},
		"Failure: Invalid params": {
			msg: types.MsgUpdateParams{
				Authority: validAuthority,
				Params: types.Params{
					TreasuryAccount: "", // invalid - empty
				},
			},
			expectedErr: types.ErrInvalidTreasuryAccount,
		},
	}
	for name, tc := range test {
		t.Run(name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.expectedErr == nil {
				require.NoError(t, err)
			} else {
				require.ErrorIs(t, err, tc.expectedErr)
			}
		})
	}
}
