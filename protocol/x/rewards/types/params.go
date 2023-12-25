package types

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/jinxprotocol/v4-chain/protocol/lib"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return Params{
		// Corresponds to module account address: jinx16wrau2x4tsg033xfrrdpae6kxfn9kyueysn4sp
		TreasuryAccount:  TreasuryAccountName,
		Denom:            lib.DefaultBaseDenom,
		DenomExponent:    lib.BaseDenomExponent,
		MarketId:         1,
		FeeMultiplierPpm: 990_000, // 0.99
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if p.TreasuryAccount == "" {
		return errorsmod.Wrap(ErrInvalidTreasuryAccount, "treasury account cannot have empty name")
	}

	if p.FeeMultiplierPpm > lib.OneMillion {
		return errorsmod.Wrap(ErrInvalidFeeMultiplierPpm, "FeeMultiplierPpm cannot be greater than 1_000_000 (100%)")
	}

	if err := sdk.ValidateDenom(p.Denom); err != nil {
		return err
	}
	return nil
}
