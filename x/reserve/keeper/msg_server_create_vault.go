package keeper

import (
	"context"

	"reserve/x/reserve/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateVault(goCtx context.Context, msg *types.MsgCreateVault) (*types.MsgCreateVaultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	collateral, err := sdk.ParseCoinNormalized(msg.Collateral)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount is not a valid Coin object")
	}
	if !collateral.IsValid() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount is not a valid Coin object")
	}

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateVaultResponse{}, nil
}
