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
	// ErrDepositNotCollateral - Deposit coin needs to be collateral coin when vault status ready
	ErrDepositNotCollateral = sdkerrors.Register(ModuleName, 3, "deposit coin not collateral") // nolint: gomnd
)
