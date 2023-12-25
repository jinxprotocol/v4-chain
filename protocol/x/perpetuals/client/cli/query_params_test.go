//go:build all || integration_test

package cli_test

import (
	"fmt"
	"testing"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	"github.com/jinxprotocol/v4-chain/protocol/x/perpetuals/client/cli"
	"github.com/jinxprotocol/v4-chain/protocol/x/perpetuals/types"
	"github.com/stretchr/testify/require"

	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
)

func TestParams(t *testing.T) {
	net, _, _ := networkWithLiquidityTierAndPerpetualObjects(t, 2, 2)
	ctx := net.Validators[0].ClientCtx

	common := []string{fmt.Sprintf("--%s=json", tmcli.OutputFlag)}

	out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdQueryParams(), common)
	require.NoError(t, err)

	var resp types.QueryParamsResponse
	require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
	require.Equal(t, types.DefaultGenesis().Params, resp.Params)
}
