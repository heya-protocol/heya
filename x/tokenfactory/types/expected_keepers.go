package types

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type BankKeeper interface {
	MintCoins(ctx context.Context, moduleAddr string, amt sdk.Coins) error
	BurnCoins(ctx context.Context, moduleAddr string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx context.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	GetBalance(ctx context.Context, addr sdk.AccAddress, denom string) sdk.Coin
	BlockedAddr(addr sdk.AccAddress) bool
	GetSupply(ctx context.Context, denom string) sdk.Coin
}

type AccountKeeper interface {
	GetModuleAddress(moduleName string) sdk.AccAddress
}
