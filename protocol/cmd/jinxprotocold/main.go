package main

import (
	"os"

	"github.com/jinxprotocol/v4-chain/protocol/app"
	"github.com/jinxprotocol/v4-chain/protocol/app/config"
	"github.com/jinxprotocol/v4-chain/protocol/cmd/jinxprotocold/cmd"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	config.SetupConfig()

	option := cmd.GetOptionWithCustomStartCmd()
	rootCmd := cmd.NewRootCmd(option)

	cmd.AddTendermintSubcommands(rootCmd)
	cmd.AddInitCmdPostRunE(rootCmd)

	if err := svrcmd.Execute(rootCmd, app.AppDaemonName, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
