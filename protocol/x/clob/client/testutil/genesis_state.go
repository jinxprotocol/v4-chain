package testutil

import (
	"testing"

	"github.com/jinxprotocol/v4-chain/protocol/testutil/constants"
	"github.com/jinxprotocol/v4-chain/protocol/testutil/network"
	satypes "github.com/jinxprotocol/v4-chain/protocol/x/subaccounts/types"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// CreateBankGenesisState returns the marshaled genesis state for the bank module.
// It will set the balance of the subaccount module in the genesis.
// If the provided subaccount module balance is negative, this function will panic.
func CreateBankGenesisState(
	t *testing.T,
	cfg network.Config,
	initialSubaccountModuleBalance int64,
) []byte {
	bankGenState := banktypes.GenesisState{
		Balances: []banktypes.Balance{
			{
				Address: satypes.ModuleAddress.String(),
				Coins: []sdk.Coin{
					sdk.NewInt64Coin(
						constants.Usdc.Denom,
						initialSubaccountModuleBalance,
					),
				},
			},
		},
	}
	bankbuf, err := cfg.Codec.MarshalJSON(&bankGenState)
	require.NoError(t, err)

	return bankbuf
}
