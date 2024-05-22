package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"reserve/x/reserve/types"
)

func (k msgServer) CreateVault(goCtx context.Context, msg *types.MsgCreateVault) (*types.MsgCreateVaultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateVaultResponse{}, nil
}
