package types_test

import (
	"testing"

	"github.com/jinxprotocol/v4-chain/protocol/x/rewards/types"
	"github.com/stretchr/testify/require"
)

func TestTreasuryModuleAddress(t *testing.T) {
	require.Equal(t, "jinx16wrau2x4tsg033xfrrdpae6kxfn9kyueysn4sp", types.TreasuryModuleAddress.String())
}
