package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/reserve module sentinel errors
var (
	// ErrDenomExists - denom already exists
	ErrMetadataExists = sdkerrors.Register(ModuleName, 1, "metadata already exists") // nolint: gomnd
)
