package tokenfactory

import (
	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"heya/x/tokenfactory/types"
)

type Keeper struct {
	storeService  store.KVStoreService
	bankKeeper    types.BankKeeper
	accountKeeper types.AccountKeeper
}

func NewKeeper(
	storeService store.KVStoreService,
	bankKeeper types.BankKeeper,
	accountKeeper types.AccountKeeper,
) Keeper {
	return Keeper{
		storeService:  storeService,
		bankKeeper:    bankKeeper,
		accountKeeper: accountKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

func (k Keeper) GetDenomAdmin(ctx sdk.Context, denom string) (string, bool) {
	store := k.storeService.OpenKVStore(ctx)
	bz, err := store.Get(types.DenomKey(denom))
	if err != nil || bz == nil {
		return "", false
	}
	return string(bz), true
}

func (k Keeper) SetDenomAdmin(ctx sdk.Context, denom, admin string) {
	store := k.storeService.OpenKVStore(ctx)
	store.Set(types.DenomKey(denom), []byte(admin))
}
