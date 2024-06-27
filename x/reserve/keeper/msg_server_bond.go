package keeper

import (
	"context"

	"reserve/x/reserve/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Bond(goCtx context.Context, msg *types.MsgBond) (*types.MsgBondResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	coin, err := sdk.ParseCoinNormalized(msg.Coin)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount is not a valid Coin object")
	}
	if !coin.IsValid() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount is not a valid Coin object")
	}

	creator, _ := sdk.AccAddressFromBech32(msg.Creator)

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, sdk.NewCoins(coin))
	if err != nil {
		return nil, err
	}

	// Bond the uid
	uid := k.GetUidCount(ctx)

	vault := types.Bond{
		Uid:   uid,
		Owner: msg.Creator,
		Denom: "ready",
		Coin:  coin,
	}

	k.SetVault(ctx, vault)

	// Update drop uid count
	k.SetUidCount(ctx, uid+1)

	return &types.MsgBondResponse{Uid: uid}, nil
}
