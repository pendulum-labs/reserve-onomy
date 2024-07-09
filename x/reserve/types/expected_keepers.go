package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	types "github.com/cosmos/cosmos-sdk/x/bank/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
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
	GetDenomMetaData(ctx sdk.Context, denom string) (types.Metadata, bool)
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	SetDenomMetaData(ctx sdk.Context, denomMetaData types.Metadata)
	GetSupply(ctx sdk.Context, denom string) sdk.Coin
	BurnCoins(ctx sdk.Context, moduleName string, amounts sdk.Coins) error
	// Methods imported from bank should be defined here
}

// MarketKeeper defines the expected interface needed to retrieve market data.
type MarketKeeper interface {
	// GetMember returns a member from its index
	GetMember(ctx sdk.Context, denomA string, denomB string) (val markettypes.Member, found bool)
	GetPool(ctx sdk.Context, pair string) (val markettypes.Pool, found bool)
	SetMember(ctx sdk.Context, member markettypes.Member)
}

// MintKeeper defines the expected interface needed to retrieve mint data
type MintKeeper interface {
	GetParams(ctx sdk.Context) (params minttypes.Params)
}
