package types_test

import (
	"testing"

	"github.com/jinxprotocol/v4-chain/protocol/x/bridge/types"
	"github.com/stretchr/testify/require"
)

func TestModuleAddress(t *testing.T) {
	require.Equal(t, "jinx1zlefkpe3g0vvm9a4h0jf9000lmqutlh94a45jv", types.ModuleAddress.String())
}
