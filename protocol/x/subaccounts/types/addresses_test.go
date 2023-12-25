package types_test

import (
	"testing"

	"github.com/jinxprotocol/v4-chain/protocol/x/subaccounts/types"
	"github.com/stretchr/testify/require"
)

func TestModuleAddress(t *testing.T) {
	require.Equal(t, "jinx1v88c3xv9xyv3eetdx0tvcmq7ung3dywp5upwc6", types.ModuleAddress.String())
}
