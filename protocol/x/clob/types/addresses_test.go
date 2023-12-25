package types_test

import (
	"testing"

	"github.com/jinxprotocol/v4-chain/protocol/x/clob/types"
	"github.com/stretchr/testify/require"
)

func TestInsuranceFundModuleAddress(t *testing.T) {
	require.Equal(t, "jinx1c7ptc87hkd54e3r7zjy92q29xkq7t79waxhcpq", types.InsuranceFundModuleAddress.String())
}
