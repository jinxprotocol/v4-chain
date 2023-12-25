package lib_test

import (
	"testing"

	"github.com/jinxprotocol/v4-chain/protocol/lib"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestMustParseCoinsNormalized(t *testing.T) {
	// Test case: valid coin string
	coinStr := "100atom"
	expectedCoins := sdk.Coins{sdk.NewInt64Coin("atom", 100)}
	result := lib.MustParseCoinsNormalized(coinStr)
	require.Equal(t, expectedCoins, result)

	// Test case: invalid coin string
	invalidCoinStr := "invalidcoin"
	require.Panics(t, func() { lib.MustParseCoinsNormalized(invalidCoinStr) })
}
