//go:build all || integration_test

package cli_test

import (
	"fmt"
	"testing"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	"github.com/jinxprotocol/v4-chain/protocol/x/clob/client/cli"
	"github.com/jinxprotocol/v4-chain/protocol/x/clob/types"
	"github.com/stretchr/testify/require"

	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
)

var (
	emptyConfig = types.BlockRateLimitConfiguration{
		MaxShortTermOrdersPerNBlocks:             []types.MaxPerNBlocksRateLimit{},
		MaxStatefulOrdersPerNBlocks:              []types.MaxPerNBlocksRateLimit{},
		MaxShortTermOrderCancellationsPerNBlocks: []types.MaxPerNBlocksRateLimit{},
	}
)

func TestCmdGetBlockRateLimitConfiguration(t *testing.T) {
	net, _ := networkWithClobPairObjects(t, 2)
	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}

	out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdGetBlockRateLimitConfiguration(), common)
	require.NoError(t, err)
	var resp types.QueryBlockRateLimitConfigurationResponse
	require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
	require.NotNil(t, resp.BlockRateLimitConfig)
	require.Equal(
		t,
		emptyConfig,
		resp.BlockRateLimitConfig,
	)
}
