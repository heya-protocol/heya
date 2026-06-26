package types

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Params struct {
	DenomCreationFee sdk.Coin
}

func DefaultParams() Params {
	return Params{
		DenomCreationFee: sdk.NewCoin("uheya", sdkmath.NewInt(1_000_000_000)),
	}
}

func (p Params) Validate() error {
	if p.DenomCreationFee.IsNil() || !p.DenomCreationFee.IsPositive() {
		return ErrInsufficientFee
	}
	return nil
}
