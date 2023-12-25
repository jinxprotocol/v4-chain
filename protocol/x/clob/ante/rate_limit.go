package ante

import (
	"github.com/jinxprotocol/v4-chain/protocol/x/clob/types"

	sdktypes "github.com/cosmos/cosmos-sdk/types"
)

var _ sdktypes.AnteDecorator = (*ClobRateLimitDecorator)(nil)

// ClobRateLimitDecorator is an AnteDecorator which is responsible for rate limiting MsgCancelOrder and MsgPlaceOrder
// requests.
//
// This AnteDecorator is a no-op if:
//   - No messages in the transaction are `MsgCancelOrder` or `MsgPlaceOrder`.
//
// This AnteDecorator returns an error if:
//   - The rate limit is exceeded for any `MsgCancelOrder` messages.
//   - The rate limit is exceeded for any `MsgPlaceOrder` messages.
//
// TODO(CLOB-721): Rate limit short term order cancellations.
type ClobRateLimitDecorator struct {
	clobKeeper types.ClobKeeper
}

func NewRateLimitDecorator(clobKeeper types.ClobKeeper) ClobRateLimitDecorator {
	return ClobRateLimitDecorator{
		clobKeeper,
	}
}

func (r ClobRateLimitDecorator) AnteHandle(
	ctx sdktypes.Context,
	tx sdktypes.Tx, simulate bool,
	next sdktypes.AnteHandler,
) (newCtx sdktypes.Context, err error) {
	for _, msg := range tx.GetMsgs() {
		switch msg := msg.(type) {
		case *types.MsgCancelOrder:
			if err = r.clobKeeper.RateLimitCancelOrder(ctx, msg); err != nil {
				return ctx, err
			}
		case *types.MsgPlaceOrder:
			if err = r.clobKeeper.RateLimitPlaceOrder(ctx, msg); err != nil {
				return ctx, err
			}
		}
	}
	return next(ctx, tx, simulate)
}
