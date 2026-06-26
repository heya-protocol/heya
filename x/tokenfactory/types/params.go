package types

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Params struct {
	DenomCreationFee sdk.Coin
}

var defaultParams = Params{
	DenomCreationFee: sdk.NewCoin("uheya", sdkmath.NewInt(1_000_000_000)),
}

func DefaultParams() Params {
	return defaultParams
}

func (p Params) Validate() error {
	if p.DenomCreationFee.IsNil() || !p.DenomCreationFee.IsPositive() {
		return ErrInsufficientFee
	}
	return nil
}
