package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"reserve/x/reserve/types"
)

func (k msgServer) Deposit(goCtx context.Context, msg *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgDepositResponse{}, nil
}
