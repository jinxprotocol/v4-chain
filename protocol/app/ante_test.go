package app_test

import (
	"testing"

	"github.com/jinxprotocol/v4-chain/protocol/app"
	testApp "github.com/jinxprotocol/v4-chain/protocol/testutil/app"
	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/x/auth/ante"
)

func newHandlerOptions() app.HandlerOptions {
	encodingConfig := app.GetEncodingConfig()
	jinxApp := testApp.DefaultTestApp(nil)
	return app.HandlerOptions{
		HandlerOptions: ante.HandlerOptions{
			AccountKeeper:   jinxApp.AccountKeeper,
			BankKeeper:      jinxApp.BankKeeper,
			SignModeHandler: encodingConfig.TxConfig.SignModeHandler(),
			FeegrantKeeper:  jinxApp.FeeGrantKeeper,
			SigGasConsumer:  ante.DefaultSigVerificationGasConsumer,
		},
		ClobKeeper: jinxApp.ClobKeeper,
	}
}

func TestNewAnteHandler(t *testing.T) {
	handlerOptions := newHandlerOptions()
	anteHandler, err := app.NewAnteHandler(handlerOptions)
	require.NoError(t, err, "NewAnteHandler call failed")
	require.NotNil(t, anteHandler, "expected non-nil AnteHandler function")
}

func TestNewAnteHandler_Error(t *testing.T) {
	tests := map[string]struct {
		handlerMutation func(*app.HandlerOptions)
		errorMsg        string
	}{
		"nil handlerOptions.AccountKeeper": {
			handlerMutation: func(options *app.HandlerOptions) { options.AccountKeeper = nil },
			errorMsg:        "account keeper is required for ante builder",
		},
		"nil handlerOptions.BankKeeper": {
			handlerMutation: func(options *app.HandlerOptions) { options.BankKeeper = nil },
			errorMsg:        "bank keeper is required for ante builder",
		},
		"nil handlerOptions.SignModeHandler": {
			handlerMutation: func(options *app.HandlerOptions) { options.SignModeHandler = nil },
			errorMsg:        "sign mode handler is required for ante builder",
		},
		"nil ClobKeeper": {
			handlerMutation: func(options *app.HandlerOptions) { options.ClobKeeper = nil },
			errorMsg:        "clob keeper is required for ante builder",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			handlerOptions := newHandlerOptions()
			tc.handlerMutation(&handlerOptions)

			anteHandler, err := app.NewAnteHandler(handlerOptions)
			require.Nil(t, anteHandler, "Expected Ante Handler creation to error")
			require.Errorf(t, err, tc.errorMsg)
		})
	}
}
