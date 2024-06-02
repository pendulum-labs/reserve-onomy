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

// GetAllVaults returns all vaults
func (k Keeper) GetVaults(ctx sdk.Context) (list []types.Vault) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VaultKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()
	var val types.Vault

	for ; iterator.Valid(); iterator.Next() {
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAllVaultsByOwner returns all vaults for a specific owner
func (k Keeper) GetAllVaultsByOwner(ctx sdk.Context, owner string) (list []types.Vault) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VaultKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	var vault types.Vault

	for ; iterator.Valid(); iterator.Next() {
		k.cdc.MustUnmarshal(iterator.Value(), &vault)
		if vault.Owner == owner {
			list = append(list, vault)
		}
	}

	return
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
func (k Keeper) SetVaultUid(ctx sdk.Context, owner string, name string, uid uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VaultMapKeyPrefix))
	vaultMapEntry := types.VaultMap{
		Uid: uid,
	}
	a := k.cdc.MustMarshal(&vaultMapEntry)
	store.Set(types.VaultMapKey(
		owner,
		name,
	), a)
}

func (k Keeper) GetVaultUid(
	ctx sdk.Context,
	owner string,
	name string,
) (uid uint64, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VaultMapKeyPrefix))

	b := store.Get(types.VaultMapKey(
		owner,
		name,
	))
	if b == nil {
		return uid, false
	}

	var vaultMapEntry types.VaultMap
	k.cdc.MustUnmarshal(b, &vaultMapEntry)
	return vaultMapEntry.Uid, true
}
