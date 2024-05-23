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

	_, found := k.GetVaultUid(ctx, msg.Creator, msg.Name)

	if found {
		return nil, sdkerrors.Wrapf(types.ErrVaultNameExists, "Vault with name %s already exists for owner", msg.Name)
	}

	creator, _ := sdk.AccAddressFromBech32(msg.Creator)

	// Use the module account as pool account
	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, sdk.NewCoins(collateral))
	if sdkError != nil {
		return nil, sdkError
	}

	// Create the uid
	uid := k.GetUidCount(ctx)

	vault := types.Vault{
		Uid:        uid,
		Owner:      msg.Creator,
		Name:       msg.Name,
		Status:     "ready",
		Collateral: collateral,
	}

	k.SetVault(ctx, vault)

	// Update drop uid count
	k.SetUidCount(ctx, uid+1)

	return &types.MsgCreateVaultResponse{Uid: uid}, nil
}
