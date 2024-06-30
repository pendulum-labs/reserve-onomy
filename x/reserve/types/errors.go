package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/reserve module sentinel errors
var (
	// ErrDenomExists - denom already exists
	ErrMetadataExists = sdkerrors.Register(ModuleName, 1, "metadata already exists") // nolint: gomnd
	// ErrDenomExists - denom already exists
	ErrVaultNameExists = sdkerrors.Register(ModuleName, 2, "vault name exists for owner") // nolint: gomnd
	// ErrVaultNotFound - vault not found
	ErrVaultNotFound = sdkerrors.Register(ModuleName, 3, "vault not found") // nolint: gomnd
	// ErrBondGtDebt - bond interest greater than debt interest
	ErrBondGtDebt = sdkerrors.Register(ModuleName, 4, "bond interest greater than debt interest") // nolint: gomnd
	// ErrInterestGtLimit - bond interest greater than debt interest
	ErrInterestGtLimit = sdkerrors.Register(ModuleName, 5, "interest rate greater than limit") // nolint: gomnd
)
