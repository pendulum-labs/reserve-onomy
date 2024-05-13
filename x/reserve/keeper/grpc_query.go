package keeper

import (
	"denom/x/reserve/types"
)

var _ types.QueryServer = Keeper{}
