package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"reserve/x/reserve/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace

		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper
		marketKeeper  types.MarketKeeper
		mintKeeper    types.MintKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,

	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	marketKeeper types.MarketKeeper,
	mintKeeper types.MintKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		paramstore:    ps,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
		marketKeeper:  marketKeeper,
		mintKeeper:    mintKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func addUid(s []uint64, r uint64) ([]uint64, bool) {
	for _, v := range s {
		if v == r {
			return s, false
		}
	}

	return append(s, r), true
}

func removeUid(s []uint64, r uint64) ([]uint64, bool) {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...), true
		}
	}
	return s, false
}

func SafeSub(intA sdk.Int, intB sdk.Int) (diff sdk.Int) {
	diff = intA.Sub(intB)
	if diff.LT(sdk.ZeroInt()) {
		panic("subtraction result negative")
	}
	return diff
}

func DebtAmount(denom types.Denom, vault types.Vault) sdk.Int {
	return (vault.DebtShares.Mul(denom.DebtDenoms)).Quo(denom.DebtShares)
}

// Collateralization Ratio = (vault_collateral * peg_rate * 100) / vault_denoms
func CollateralizationRatio(denom types.Denom, vault types.Vault, rate sdk.Dec) sdk.Int {
	return (((vault.Collateral.Amount).Mul(sdk.NewInt(100))).ToDec().Mul(rate).Quo(DebtAmount(denom, vault).ToDec())).RoundInt()
}
