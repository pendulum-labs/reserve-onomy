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

// GetAllDenom returns all denoms
func (k Keeper) GetAllDenom(ctx sdk.Context) (list []types.Denom) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DenomKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Denom
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
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

// GetDenom returns a denom from its index
func (k Keeper) ApplyInterest(ctx sdk.Context) {
	blocksPerYear := k.mintKeeper.GetParams(ctx).BlocksPerYear
	denoms := k.GetAllDenom(ctx)

	for _, denom := range denoms {
		debtYearlyProvisions := (sdk.NewIntFromUint64(denom.DebtInterestRate).Mul(denom.DebtDenoms)).Quo(sdk.NewInt(10000))
		denom.DebtDenoms = denom.DebtDenoms.Add(debtYearlyProvisions.Quo(sdk.NewIntFromUint64(blocksPerYear)))
		bondYearlyProvisions := (sdk.NewIntFromUint64(denom.BondInterestRate).Mul(denom.BondDenoms)).Quo(sdk.NewInt(10000))
		denom.BondDenoms = denom.BondDenoms.Add(bondYearlyProvisions.Quo(sdk.NewIntFromUint64(blocksPerYear)))
		k.SetDenom(ctx, denom)
	}

}
