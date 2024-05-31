package keeper

import (
	"reserve/x/reserve/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetReserve set a specific reserve in the store from its index
func (k Keeper) SetReserve(ctx sdk.Context, reserve types.Reserve) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ReserveKey)
	a := k.cdc.MustMarshal(&reserve)
	store.Set(byteKey, a)
}

// GetReserve returns a reserve from its index
func (k Keeper) GetReserve(ctx sdk.Context) (val types.Reserve, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ReserveKey)
	b := store.Get(byteKey)
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}
