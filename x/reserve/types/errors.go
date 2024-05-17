package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/reserve module sentinel errors
var (
	// ErrDenomExists - denom already exists
	ErrDenomExists = sdkerrors.Register(ModuleName, 1, "denom already exists") // nolint: gomnd
)
