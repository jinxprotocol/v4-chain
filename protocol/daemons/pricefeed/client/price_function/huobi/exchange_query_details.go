package huobi

import (
	"github.com/jinxprotocol/v4-chain/protocol/daemons/pricefeed/client/constants/exchange_common"
	"github.com/jinxprotocol/v4-chain/protocol/daemons/pricefeed/client/types"
)

var (
	HuobiDetails = types.ExchangeQueryDetails{
		Exchange:      exchange_common.EXCHANGE_ID_HUOBI,
		Url:           "https://api.huobi.pro/market/tickers",
		PriceFunction: HuobiPriceFunction,
		IsMultiMarket: true,
	}
)
