package app_test

import (
	"testing"

	"github.com/jinxprotocol/v4-chain/protocol/testutil/app"
	"github.com/stretchr/testify/require"
)

func TestExportAppStateAndValidators_Panics(t *testing.T) {
	jinxApp := app.DefaultTestApp(nil)
	require.Panics(t, func() { jinxApp.ExportAppStateAndValidators(false, nil, nil) }) // nolint:errcheck
}
