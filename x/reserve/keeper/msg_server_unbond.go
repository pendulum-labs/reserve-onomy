package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"reserve/x/reserve/types"
)

func (k msgServer) Unbond(goCtx context.Context, msg *types.MsgUnbond) (*types.MsgUnbondResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUnbondResponse{}, nil
}
