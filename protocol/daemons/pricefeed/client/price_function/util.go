package price_function

import (
	"errors"
	"fmt"
	"math/big"
	"regexp"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/jinxprotocol/v4-chain/protocol/daemons/pricefeed/types"
	"github.com/jinxprotocol/v4-chain/protocol/lib"
)

var (
	apiResponseValidator *validator.Validate

	// exchangeSideErrorResponsePatterns captures a list of error strings patterns are generated by exchange APIs.
	// These errors are not attributable to the exchange by http status code, but rather by the response body.
	exchangeSideErrorResponsePatterns = []string{
		// An error that occurs server-side on the exchange. These are generic error strings that are
		// percolated up by go's http2 library.
		// nolint:lll
		// See https://stackoverflow.com/questions/45209168/http2-server-sent-goaway-and-closed-the-connection-laststreamid-1999
		`http2: server sent GOAWAY and closed the connection`,

		// http2 error: Server closed connection. Could sometimes be due to rate limiting, as not all exchanges return
		// 429s.
		"http2: client connection force closed via ClientConn.Close",

		// This can sometimes occur with TLS errors, or could arise as a result of gzip decompression by the go
		// http2 library.
		"unexpected EOF",

		// Existing error response identified as an internal server error by status code.
		"Unexpected response status code of: (\\d+)",

		"Failed to get entities from store",

		"request timeout",

		// "internal error", "Internal error", "INTERNAL_ERROR", "SYS_ERROR"
		"([Ii]nternal error|INTERNAL_ERROR|SYS_ERROR)",
	}
)

// IsGenericExchangeError returns true if the error message has been positively identified as being due to an exchange
// side error. These errors could occur on any exchange.
func IsGenericExchangeError(err error) bool {
	for _, exchangeSideErrorResponse := range exchangeSideErrorResponsePatterns {
		if match, err := regexp.Match(exchangeSideErrorResponse, []byte(err.Error())); err == nil && match {
			return true
		}
	}
	return false
}

// validatePositiveNumericString is a custom validation function that ensures a particular string field in
// a struct being validated can be parsed into a positive-valued float. We register this function in order
// to ensure that returned numeric string values in the Kraken response do not represent zero or negative numbers.
// To see where this is used, note the `validate:"positive-float-string"` struct tag in the KrakenTickerResult.
func validatePositiveNumericString(fl validator.FieldLevel) bool {
	val, err := strconv.ParseFloat(fl.Field().String(), 64)
	if err != nil {
		return false
	}
	return val > 0
}

// GetApiResponseValidator returns a validator with custom logic registered to validate fields returned by
// various exchange API responses.
func GetApiResponseValidator() (*validator.Validate, error) {
	if apiResponseValidator == nil {
		validate := validator.New()
		err := validate.RegisterValidation("positive-float-string", validatePositiveNumericString)
		if err != nil {
			return nil, fmt.Errorf("API response validation internal error (%w)", err)
		}
		apiResponseValidator = validate
	}
	return apiResponseValidator, nil
}

// GetOnlyTickerAndExponent returns the only ticker and exponent in the provided
// `tickerToExponent` map. If the map contains more than one key, an error is returned.
func GetOnlyTickerAndExponent(
	tickerToExponent map[string]int32,
	exchange string,
) (string, int32, error) {
	// Verify exactly one market is expected from the response.
	if len(tickerToExponent) != 1 {
		return "", 0, fmt.Errorf(
			"Invalid market price exponent map for %v price function of length: %v, expected length 1",
			exchange,
			len(tickerToExponent),
		)
	}

	// Get ticker and value of exponent.
	var ticker string
	var exponent int32
	// Set `ticker` and `exponent` explicitly from the for loop.
	// tickerToExponent has only one entry so the for loop only runs once.
	for ticker, exponent = range tickerToExponent {
	}

	return ticker, exponent, nil
}

// GetUint64MedianFromReverseShiftedBigFloatValues shifts all values in a slice of floats by an
// exponent, converts the shifted slice values to uint64 and then returns the median of the slice.
// 1) Verify length of slice is > 0.
// 2) Reverse shift big float price values by the exponent for the market.
// 3) Convert big float values to uint64 values.
// 4) Get the median value from the uint64 price values and return.
func GetUint64MedianFromReverseShiftedBigFloatValues(
	bigFloatSlice []*big.Float,
	exponent int32,
	resolver func([]uint64) (uint64, error),
) (uint64, error) {
	// 1) Verify length of slice is > 0.
	if len(bigFloatSlice) == 0 {
		return 0, errors.New("Invalid input: big float slice must contain values to medianize")
	}

	// 2) Reverse shift big float price values by the exponent for the market.
	updatedBigFloatSlice := reverseShiftBigFloatSlice(bigFloatSlice, exponent)

	// 3) Convert big float values to uint64 values.
	uint64Slice, err := lib.ConvertBigFloatSliceToUint64Slice(updatedBigFloatSlice)
	if err != nil {
		return 0, err
	}

	// 4) Get the median value from the uint64 price values and return.
	return resolver(uint64Slice)
}

// reverseShiftBigFloatSlice shifts the given floats by exponent in the reverse direction.
// If the exponent is 0, then do nothing (i.e. `123.456` -> `123.456`)
// If the exponent is positive, then shift to the right (i.e. exponent = 1, `123.456` -> `12.3456`)
// If the exponent is negative, then shift to the left (i.e. exponent = -1, `123.456` -> `1234.56`)
func reverseShiftBigFloatSlice(
	values []*big.Float,
	exponent int32,
) []*big.Float {
	unsignedExponent := lib.AbsInt32(exponent)

	pow10 := new(big.Float).SetInt(lib.BigPow10(uint64(unsignedExponent)))
	updatedValues := make([]*big.Float, 0, len(values))
	for _, value := range values {
		updatedValues = append(updatedValues, reverseShiftFloatWithPow10(value, pow10, exponent))
	}
	return updatedValues
}

func reverseShiftFloatWithPow10(value *big.Float, pow10 *big.Float, exponent int32) *big.Float {
	if exponent == 0 {
		return value
	} else if exponent > 0 {
		return new(big.Float).Quo(value, pow10)
	} else { // exponent < 0
		return new(big.Float).Mul(value, pow10)
	}
}

// Ticker encodes a ticker response returned by an exchange API. It contains accessors for the ticker's
// ask price, bid price, and last price, which are medianized to compute an exchange price.
type Ticker interface {
	GetPair() string
	GetAskPrice() string
	GetBidPrice() string
	GetLastPrice() string
}

// GetMedianPricesFromTickers processes through a list of `tickers` and calculates a median price (from
// `LastPrice`, `AskPrice`, and `BidPrice`) for each ticker in `tickerToExponent` and marks a ticker
// as unavailable if it's not present in `tickers` or its ticker's validation or calculation fails.
func GetMedianPricesFromTickers[T Ticker](
	tickers []T,
	tickerToExponent map[string]int32,
	resolver types.Resolver,
) (tickerToPrice map[string]uint64, unavailableTickers map[string]error, err error) {
	// Create API response validator, if not already.
	if apiResponseValidator == nil {
		apiResponseValidator, err = GetApiResponseValidator()
		if err != nil {
			return nil, nil, fmt.Errorf("Error creating API response validator (%w)", err)
		}
	}

	tickerToPrice = make(map[string]uint64, len(tickerToExponent))
	unavailableTickers = make(map[string]error)

	// Iterate through every ticker in response and calculate median prices for requested
	// tickers (the ones present in `tickerToExponent`).
	for _, ticker := range tickers {
		tickerPair := ticker.GetPair()
		if exponent, exists := tickerToExponent[tickerPair]; exists {
			// Validate ticker info.
			err = apiResponseValidator.Struct(ticker)
			// If validation failed, mark tickerPair as unavailable with validation error.
			if err != nil {
				unavailableTickers[tickerPair] = err
				continue
			}
			// Get big float values from prices in response.
			bigFloatSlice, err := lib.ConvertStringSliceToBigFloatSlice(
				[]string{ticker.GetAskPrice(), ticker.GetBidPrice(), ticker.GetLastPrice()})
			// If unsuccessful, mark ticker as unavailable with corresponding error.
			if err != nil {
				unavailableTickers[tickerPair] = err
				continue
			}

			// Get the median uint64 value from the slice of big float price values.
			medianPrice, err := GetUint64MedianFromReverseShiftedBigFloatValues(
				bigFloatSlice,
				exponent,
				resolver,
			)
			// If unsuccessful, mark tickerPair as unavailable with corresponding error.
			// Otherwise, store median price for corresponding tickerPair.
			if err != nil {
				unavailableTickers[tickerPair] = err
				continue
			} else {
				tickerToPrice[tickerPair] = medianPrice
			}
		}
	}

	// Iterate through every requested ticker and mark as unavailable if it wasn't
	// seen / processed in above loop.
	for ticker := range tickerToExponent {
		_, priceCalculationSucceeded := tickerToPrice[ticker]
		_, priceCalculationErrored := unavailableTickers[ticker]
		if !priceCalculationSucceeded && !priceCalculationErrored {
			unavailableTickers[ticker] = fmt.Errorf("no listing found for ticker %v", ticker)
		}
	}

	return tickerToPrice, unavailableTickers, nil
}

// ConvertFloat64ToString converts a `float64` to `string`.
func ConvertFloat64ToString(num float64) string {
	return strconv.FormatFloat(num, 'f', -1, 64)
}
