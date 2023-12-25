package types_test

import (
	"testing"

	"github.com/jinxprotocol/v4-chain/protocol/x/delaymsg/types"
	"github.com/stretchr/testify/require"
)

func TestModuleAddress(t *testing.T) {
	require.Equal(t, "jinx1mkkvp26dngu6n8rmalaxyp3gwkjuzztqn3pafr", types.ModuleAddress.String())
}
