package cli

import (
	"context"

	"github.com/jinxprotocol/v4-chain/protocol/x/clob/types"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
)

func CmdGetEquityTierLimitConfig() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-equity-tier-limit-config",
		Short: "get the equity tier limit configuration",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryEquityTierLimitConfigurationRequest{}

			res, err := queryClient.EquityTierLimitConfiguration(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
