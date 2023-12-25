package cmd

import (
	"errors"
	"io"
	"os"
	"path/filepath"

	rosettaCmd "cosmossdk.io/tools/rosetta/cmd"

	dbm "github.com/cometbft/cometbft-db"
	tmcli "github.com/cometbft/cometbft/libs/cli"
	"github.com/cometbft/cometbft/libs/log"
	tmtypes "github.com/cometbft/cometbft/types"
	jinxapp "github.com/jinxprotocol/v4-chain/protocol/app"
	"github.com/jinxprotocol/v4-chain/protocol/app/basic_manager"
	protocolflags "github.com/jinxprotocol/v4-chain/protocol/app/flags"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/config"
	"github.com/cosmos/cosmos-sdk/client/debug"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/server"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/snapshots"
	snapshottypes "github.com/cosmos/cosmos-sdk/snapshots/types"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"

	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	// Unnamed import of statik for swagger UI support.
	// Used in cosmos-sdk when registering the route for swagger docs.
	_ "github.com/jinxprotocol/v4-chain/protocol/client/docs/statik"
)

const (
	EnvPrefix = "JINX"

	flagIAVLCacheSize = "iavl-cache-size"
)

// TODO(DEC-1097): improve `cmd/` by adding tests, custom app configs, custom init cmd, and etc.
// NewRootCmd creates a new root command for `jinxprotocold`. It is called once in the main function.
func NewRootCmd(
	option *RootCmdOption,
) *cobra.Command {
	return NewRootCmdWithInterceptors(
		option,
		func(serverCtxPtr *server.Context) {

		},
		func(s string, appConfig *JinxAppConfig) (string, *JinxAppConfig) {
			return s, appConfig
		},
		func(app *jinxapp.App) *jinxapp.App {
			return app
		},
	)
}

func NewRootCmdWithInterceptors(
	option *RootCmdOption,
	serverCtxInterceptor func(serverCtxPtr *server.Context),
	appConfigInterceptor func(string, *JinxAppConfig) (string, *JinxAppConfig),
	appInterceptor func(app *jinxapp.App) *jinxapp.App,
) *cobra.Command {
	encodingConfig := jinxapp.GetEncodingConfig()
	initClientCtx := client.Context{}.
		WithCodec(encodingConfig.Codec).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithInput(os.Stdin).
		WithAccountRetriever(types.AccountRetriever{}).
		WithBroadcastMode(flags.BroadcastSync).
		WithHomeDir(jinxapp.DefaultNodeHome).
		WithViper(EnvPrefix)

	rootCmd := &cobra.Command{
		Use:   jinxapp.AppDaemonName,
		Short: "Start jinxprotocol app",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// set the default command outputs
			cmd.SetOut(cmd.OutOrStdout())
			cmd.SetErr(cmd.ErrOrStderr())

			initClientCtx, err := client.ReadPersistentCommandFlags(initClientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			initClientCtx, err = config.ReadFromClientConfig(initClientCtx)
			if err != nil {
				return err
			}

			if err := client.SetCmdClientContextHandler(initClientCtx, cmd); err != nil {
				return err
			}

			customAppTemplate, customAppConfig := appConfigInterceptor(initAppConfig())
			customTMConfig := initTendermintConfig()

			if err := server.InterceptConfigsPreRunHandler(
				cmd,
				customAppTemplate,
				customAppConfig,
				customTMConfig,
			); err != nil {
				return err
			}

			serverCtx := server.GetServerContextFromCmd(cmd)

			// Format logs for error tracking if it is enabled via flags.
			if ddErrorTrackingFormatterEnabled :=
				serverCtx.Viper.Get(protocolflags.DdErrorTrackingFormat); ddErrorTrackingFormatterEnabled != nil {
				if enabled, err := cast.ToBoolE(ddErrorTrackingFormatterEnabled); err == nil && enabled {
					jinxapp.SetZerologDatadogErrorTrackingFormat()
				}
			}
			serverCtxInterceptor(serverCtx)

			return nil
		},
		SilenceUsage: true,
	}

	initRootCmd(rootCmd, option, encodingConfig, appInterceptor)

	return rootCmd
}

// initRootCmd initializes the app's root command with useful commands.
func initRootCmd(
	rootCmd *cobra.Command,
	option *RootCmdOption,
	encodingConfig jinxapp.EncodingConfig,
	appInterceptor func(app *jinxapp.App) *jinxapp.App,
) {
	gentxModule := basic_manager.ModuleBasics[genutiltypes.ModuleName].(genutil.AppModuleBasic)
	rootCmd.AddCommand(
		genutilcli.InitCmd(basic_manager.ModuleBasics, jinxapp.DefaultNodeHome),
		genutilcli.CollectGenTxsCmd(banktypes.GenesisBalancesIterator{}, jinxapp.DefaultNodeHome, gentxModule.GenTxValidator),
		genutilcli.MigrateGenesisCmd(),
		genutilcli.GenTxCmd(
			basic_manager.ModuleBasics,
			encodingConfig.TxConfig,
			banktypes.GenesisBalancesIterator{},
			jinxapp.DefaultNodeHome,
		),
		genutilcli.ValidateGenesisCmd(basic_manager.ModuleBasics),
		AddGenesisAccountCmd(jinxapp.DefaultNodeHome),
		tmcli.NewCompletionCmd(rootCmd, true),
		debug.Cmd(),
		config.Cmd(),
	)

	a := appCreator{encodingConfig}
	server.AddCommands(
		rootCmd,
		jinxapp.DefaultNodeHome,
		func(logger log.Logger, db dbm.DB, writer io.Writer, options servertypes.AppOptions) servertypes.Application {
			return appInterceptor(a.newApp(logger, db, writer, options))
		},
		a.appExport,
		func(cmd *cobra.Command) {
			addModuleInitFlags(cmd)

			if option.startCmdCustomizer != nil {
				option.startCmdCustomizer(cmd)
			}
		},
	)

	// add keybase, auxiliary RPC, query, and tx child commands
	rootCmd.AddCommand(
		rpc.StatusCommand(),
		queryCommand(),
		txCommand(),
		keys.Commands(jinxapp.DefaultNodeHome),
	)

	// add rosetta commands
	rootCmd.AddCommand(rosettaCmd.RosettaCommand(encodingConfig.InterfaceRegistry, encodingConfig.Codec))
}

// addModuleInitFlags adds module specific init flags.
func addModuleInitFlags(startCmd *cobra.Command) {
	crisis.AddModuleInitFlags(startCmd)
}

// queryCommand adds transaction and account querying commands.
func queryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "query",
		Aliases:                    []string{"q"},
		Short:                      "Querying subcommands",
		DisableFlagParsing:         false,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		authcmd.GetAccountCmd(),
		rpc.ValidatorCommand(),
		rpc.BlockCommand(),
		authcmd.QueryTxsByEventsCmd(),
		authcmd.QueryTxCmd(),
	)

	basic_manager.ModuleBasics.AddQueryCommands(cmd)
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}

// txCommand adds transaction signing, encoding / decoding, and broadcasting commands.
func txCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "tx",
		Short:                      "Transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		authcmd.GetSignCommand(),
		authcmd.GetSignBatchCommand(),
		authcmd.GetMultiSignCommand(),
		authcmd.GetMultiSignBatchCmd(),
		authcmd.GetValidateSignaturesCommand(),
		flags.LineBreak,
		authcmd.GetBroadcastCommand(),
		authcmd.GetEncodeCommand(),
		authcmd.GetDecodeCommand(),
	)

	basic_manager.ModuleBasics.AddTxCommands(cmd)
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}

type appCreator struct {
	encCfg jinxapp.EncodingConfig
}

// newApp initializes and returns a new app.
func (a appCreator) newApp(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	appOpts servertypes.AppOptions,
) *jinxapp.App {
	var cache sdk.MultiStorePersistentCache

	if cast.ToBool(appOpts.Get(server.FlagInterBlockCache)) {
		cache = store.NewCommitKVStoreCacheManager()
	}

	skipUpgradeHeights := make(map[int64]bool)
	for _, h := range cast.ToIntSlice(appOpts.Get(server.FlagUnsafeSkipUpgrades)) {
		skipUpgradeHeights[int64(h)] = true
	}

	pruningOpts, err := server.GetPruningOptionsFromFlags(appOpts)
	if err != nil {
		panic(err)
	}

	homeDir := cast.ToString(appOpts.Get(flags.FlagHome))
	chainID := cast.ToString(appOpts.Get(flags.FlagChainID))
	if chainID == "" {
		// fallback to genesis chain-id
		appGenesis, err := tmtypes.GenesisDocFromFile(filepath.Join(homeDir, "config", "genesis.json"))
		if err != nil {
			panic(err)
		}

		chainID = appGenesis.ChainID
	}

	snapshotDir := filepath.Join(cast.ToString(appOpts.Get(flags.FlagHome)), "data", "snapshots")
	snapshotDB, err := dbm.NewDB("metadata", server.GetAppDBBackend(appOpts), snapshotDir)
	if err != nil {
		panic(err)
	}
	snapshotStore, err := snapshots.NewStore(snapshotDB, snapshotDir)
	if err != nil {
		panic(err)
	}

	snapshotOptions := snapshottypes.NewSnapshotOptions(
		cast.ToUint64(appOpts.Get(server.FlagStateSyncSnapshotInterval)),
		cast.ToUint32(appOpts.Get(server.FlagStateSyncSnapshotKeepRecent)),
	)

	return jinxapp.New(
		logger,
		db,
		snapshotDB,
		traceStore,
		true,
		appOpts,
		baseapp.SetPruning(pruningOpts),
		baseapp.SetMinGasPrices(cast.ToString(appOpts.Get(server.FlagMinGasPrices))),
		baseapp.SetHaltHeight(cast.ToUint64(appOpts.Get(server.FlagHaltHeight))),
		baseapp.SetHaltTime(cast.ToUint64(appOpts.Get(server.FlagHaltTime))),
		baseapp.SetMinRetainBlocks(cast.ToUint64(appOpts.Get(server.FlagMinRetainBlocks))),
		baseapp.SetInterBlockCache(cache),
		baseapp.SetTrace(cast.ToBool(appOpts.Get(server.FlagTrace))),
		baseapp.SetIndexEvents(cast.ToStringSlice(appOpts.Get(server.FlagIndexEvents))),
		baseapp.SetSnapshot(snapshotStore, snapshotOptions),
		baseapp.SetIAVLCacheSize(int(cast.ToUint64(appOpts.Get(flagIAVLCacheSize)))),
		baseapp.SetChainID(chainID),
	)
}

// appExport creates and exports a new app, returns the state of the app for a genesis file.
//
// Deprecated: this feature relies on the use of known unstable, legacy export functionality
// from cosmos.
func (a appCreator) appExport(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	height int64,
	forZeroHeight bool,
	jailAllowedAddrs []string,
	appOpts servertypes.AppOptions,
	modulesToExport []string,
) (servertypes.ExportedApp, error) {
	homePath, ok := appOpts.Get(flags.FlagHome).(string)
	if !ok || homePath == "" {
		return servertypes.ExportedApp{}, errors.New("application home not set")
	}

	jinxApp := jinxapp.New(
		logger,
		db,
		dbm.NewMemDB(),
		traceStore,
		height == -1, // -1: no height provided
		appOpts,
	)

	if height != -1 {
		if err := jinxApp.LoadHeight(height); err != nil {
			return servertypes.ExportedApp{}, err
		}
	}

	return jinxApp.ExportAppStateAndValidators(forZeroHeight, jailAllowedAddrs, modulesToExport)
}
