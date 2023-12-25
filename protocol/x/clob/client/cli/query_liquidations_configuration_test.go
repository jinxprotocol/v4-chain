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

func TestCmdGetLiquidationsConfiguration(t *testing.T) {
	net, _ := networkWithClobPairObjects(t, 2)
	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}

	out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdGetLiquidationsConfiguration(), common)
	require.NoError(t, err)
	var resp types.QueryLiquidationsConfigurationResponse
	require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
	require.NotNil(t, resp.LiquidationsConfig)
	require.Equal(
		t,
		types.LiquidationsConfig_Default,
		resp.LiquidationsConfig,
	)
}
