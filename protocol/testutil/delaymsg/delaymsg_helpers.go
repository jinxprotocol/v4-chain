package delaymsg

import (
	"testing"

	"github.com/jinxprotocol/v4-chain/protocol/testutil/constants"
	"github.com/stretchr/testify/require"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
)

// CreateTestAnyMsg returns an encoded Any object for an sdk.Msg for testing. This is useful
// when a valid message is needed, but the message will never be executed.
func CreateTestAnyMsg(t *testing.T) *codectypes.Any {
	any, err := codectypes.NewAnyWithValue(constants.TestMsg1)
	require.NoError(t, err)
	return any
}
