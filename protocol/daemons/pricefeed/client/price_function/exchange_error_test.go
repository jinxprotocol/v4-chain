package price_function_test

import (
	"errors"
	"testing"

	"github.com/jinxprotocol/v4-chain/protocol/daemons/pricefeed/client/price_function"
	"github.com/stretchr/testify/require"
)

func TestExchangeError(t *testing.T) {
	error := price_function.NewExchangeError("exchange", "error")
	var exchangeError price_function.ExchangeError
	found := errors.As(error, &exchangeError)
	require.True(t, found)
	require.Equal(t, error, exchangeError)
	require.Equal(t, "exchange", exchangeError.GetExchangeId())
}
