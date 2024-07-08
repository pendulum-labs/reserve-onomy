package keeper

import (
	"reserve/x/reserve/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetDenom set a specific denom in the store from its index
func (k Keeper) SetDenom(ctx sdk.Context, denom types.Denom) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DenomKeyPrefix))
	a := k.cdc.MustMarshal(&denom)
	store.Set(types.DenomKey(
		denom.DenomBase,
	), a)
}

// GetDenom returns a denom from its index
func (k Keeper) GetDenom(
	ctx sdk.Context,
	base string,
) (denom types.Denom, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DenomKeyPrefix))

	b := store.Get(types.DenomKey(
		base,
	))
	if b == nil {
		return denom, false
	}

	k.cdc.MustUnmarshal(b, &denom)
	return denom, true
}

// RemoveDenom removes a denom from the store
func (k Keeper) RemoveDenom(
	ctx sdk.Context,
	base string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DenomKeyPrefix))

	b := store.Get(types.DenomKey(
		base,
	))

	if b == nil {
		return
	}

	store.Delete(types.DenomKey(
		base,
	))
}

// SetDenom set a specific denom in the store from its index
func (k Keeper) SetBonded(ctx sdk.Context, bonded types.Bonded) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DenomKeyPrefix))
	a := k.cdc.MustMarshal(&bonded)
	store.Set(types.DenomKey(
		bonded.BondBase,
	), a)
}

// GetDenom returns a denom from its index
func (k Keeper) GetBonded(
	ctx sdk.Context,
	bondBase string,
) (bonded types.Bonded, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DenomKeyPrefix))

	b := store.Get(types.DenomKey(
		bondBase,
	))
	if b == nil {
		return bonded, false
	}

	k.cdc.MustUnmarshal(b, &bonded)
	return bonded, true
}
