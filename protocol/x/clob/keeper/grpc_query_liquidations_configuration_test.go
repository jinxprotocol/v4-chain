package keeper_test

import (
	"testing"

	testApp "github.com/jinxprotocol/v4-chain/protocol/testutil/app"
	"github.com/jinxprotocol/v4-chain/protocol/x/clob/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdktypes "github.com/cosmos/cosmos-sdk/types"
)

func TestLiquidationsConfiguration(t *testing.T) {
	tests := map[string]struct {
		req *types.QueryLiquidationsConfigurationRequest
		res *types.QueryLiquidationsConfigurationResponse
		err error
	}{
		"success": {
			req: &types.QueryLiquidationsConfigurationRequest{},
			res: &types.QueryLiquidationsConfigurationResponse{
				LiquidationsConfig: types.LiquidationsConfig_Default,
			},
		},
		"failure: nil request": {
			req: nil,
			err: status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tApp := testApp.NewTestAppBuilder(t).Build()
			ctx := tApp.InitChain()
			res, err := tApp.App.ClobKeeper.LiquidationsConfiguration(sdktypes.WrapSDKContext(ctx), tc.req)

			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.res, res)
			}
		})
	}
}
