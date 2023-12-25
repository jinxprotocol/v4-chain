package lib_test

import (
	"testing"

	"github.com/jinxprotocol/v4-chain/protocol/lib"
	"github.com/stretchr/testify/require"
)

func TestGovModuleAddress(t *testing.T) {
	require.Equal(t, "jinx10d07y265gmmuvt4z0w9aw880jnsr700j5g6x5y", lib.GovModuleAddress.String())
}
