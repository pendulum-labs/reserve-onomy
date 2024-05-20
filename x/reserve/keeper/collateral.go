package keeper

import (
	"reserve/x/reserve/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetCollateral set a specific collateral in the store from its index
func (k Keeper) SetCollateral(ctx sdk.Context, collateral types.Collateral) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CollateralKeyPrefix))
	a := k.cdc.MustMarshal(&collateral)
	store.Set(types.CollateralKey(
		collateral.Base,
	), a)
}

// GetCollateral returns a collateral from its index
func (k Keeper) GetCollateral(
	ctx sdk.Context,
	uid uint64,
) (val types.Collateral, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CollateralKeyPrefix))

	b := store.Get(types.CollateralKey(
		val.Base,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCollateral removes a collateral from the store
func (k Keeper) RemoveCollateral(
	ctx sdk.Context,
	base string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CollateralKeyPrefix))

	b := store.Get(types.CollateralKey(
		base,
	))

	if b == nil {
		return
	}

	store.Delete(types.CollateralKey(
		base,
	))
}
