package keeper

import (
	"reserve/x/reserve/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetVault set a specific vault in the store from its index
func (k Keeper) SetVault(ctx sdk.Context, vault types.Vault) {
	store1 := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VaultKeyPrefix))
	a := k.cdc.MustMarshal(&vault)
	store1.Set(types.VaultKey(
		vault.Uid,
	), a)

	store2 := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VaultsKeyPrefix))

	b := store2.Get(types.VaultsKey(
		vault.Owner,
	))
	var vaults types.Vaults

	if b == nil {
		vaults = types.Vaults{
			Owner: vault.Owner,
			Uids:  []uint64{vault.Uid},
		}
	} else {
		k.cdc.MustUnmarshal(b, &vaults)

		if !contains(vaults.Uids, vault.Uid) {
			vaults.Uids = append(vaults.Uids, vault.Uid)
		}
	}

	c := k.cdc.MustMarshal(&vaults)
	store1.Set(types.VaultsKey(
		vault.Owner,
	), c)
}

func contains(s []uint64, e uint64) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// GetVault returns a vault from its index
func (k Keeper) GetVault(
	ctx sdk.Context,
	uid uint64,
) (val types.Vault, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VaultKeyPrefix))

	b := store.Get(types.VaultKey(
		uid,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveVault removes a vault from the store
func (k Keeper) RemoveVault(
	ctx sdk.Context,
	uid uint64,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VaultKeyPrefix))

	b := store.Get(types.VaultKey(
		uid,
	))

	if b == nil {
		return
	}

	store.Delete(types.VaultKey(
		uid,
	))
}

// SetVault set a specific vault in the store from its index
func (k Keeper) SetVaults(ctx sdk.Context, vaults types.Vaults) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VaultKeyPrefix))
	a := k.cdc.MustMarshal(&vaults)
	store.Set(types.VaultsKey(
		vaults.Owner,
	), a)
}

// GetVault returns a vault from its index
func (k Keeper) GetVaults(
	ctx sdk.Context,
	owner string,
) (val types.Vaults, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VaultKeyPrefix))

	b := store.Get(types.VaultsKey(
		owner,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// VaultNameExists
func (k Keeper) VaultNameExists(
	ctx sdk.Context,
	owner string,
	name string,
) (exists bool) {
	vaults, found := k.GetVaults(ctx, owner)

	if found {
		var vault types.Vault
		for _, uid := range vaults.Uids {
			vault, _ = k.GetVault(ctx, uid)
			if vault.Name == name {
				return true
			}
		}
	}

	return false
}
