package metrics

import (
	gometrics "github.com/armon/go-metrics"
	"github.com/jinxprotocol/v4-chain/protocol/daemons/pricefeed/client/types"
	"github.com/jinxprotocol/v4-chain/protocol/lib/metrics"
)

const (
	INVALID = "INVALID"
)

// GetLabelForMarketId converts a marketId uint32 into a human-readable market pair and then returns the
// label with the market pair.
func GetLabelForMarketId(marketId types.MarketId) gometrics.Label {
	marketPair := GetMarketPairForTelemetry(marketId)
	return metrics.GetLabelForStringValue(metrics.MarketId, marketPair)
}

// GetLabelForExchangeId converts an exchangeId uint32 into a name and then
// returns the label with the name.
func GetLabelForExchangeId(exchangeId types.ExchangeId) gometrics.Label {
	return metrics.GetLabelForStringValue(metrics.ExchangeId, exchangeId)
}
