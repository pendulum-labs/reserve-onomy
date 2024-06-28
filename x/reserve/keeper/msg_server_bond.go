package keeper

import (
	"context"

	"reserve/x/reserve/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Bond(goCtx context.Context, msg *types.MsgBond) (*types.MsgBondResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	denom, err := sdk.ParseCoinNormalized(msg.Denom)
	if err != nil {
		return nil, err
	}

	_ = ctx

	return &types.MsgBondResponse{}, nil
}
