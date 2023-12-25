package abci

import (
	"fmt"

	"github.com/pkg/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RunCached wraps a function with a cache context and writes the cache to the context if the
// function call succeeds. If the function call fails, the cache is discarded.
func RunCached(
	c sdk.Context,
	f func(sdk.Context) error,
) (
	err error,
) {
	ctx, writeCache := c.CacheContext()

	defer func() {
		if r := recover(); r != nil {
			err = errors.WithStack(fmt.Errorf("recovered from panic in cached context: %v", r))
		}
	}()

	if err := f(ctx); err != nil {
		return err
	}

	writeCache()
	return nil
}
