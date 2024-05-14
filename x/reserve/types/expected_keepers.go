package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	markettypes "github.com/pendulum-labs/market/x/market/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}

// MarketKeeper defines the expected interface needed to retrieve market data.
type MarketKeeper interface {
	// GetMember returns a member from its index
	GetMember(ctx sdk.Context, denomA string, denomB string) (val markettypes.Member, found bool)
	GetPool(ctx sdk.Context, pair string) (val markettypes.Pool, found bool)
	SetMember(ctx sdk.Context, member markettypes.Member)
}
