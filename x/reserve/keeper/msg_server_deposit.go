package keeper

import (
	"context"
	"strconv"

	"reserve/x/reserve/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Deposit(goCtx context.Context, msg *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	uid, err := strconv.ParseUint(msg.Uid, 10, 64)
	if err != nil {
		return nil, err
	}

	coin, err := sdk.ParseCoinNormalized(msg.Coin)
	if err != nil {
		return nil, err
	}

	vault, found := k.GetVault(ctx, uid)
	if found {
		switch {
		case vault.Status == "ready" && vault.Collateral.Denom == coin.Denom:
			vault.Collateral = vault.Collateral.Add(coin)
		case vault.Status == "active" && vault.Collateral.Denom == coin.Denom:
			vault.Collateral = vault.Collateral.Add(coin)
		case vault.Status == "active" && vault.Denom.Denom == coin.Denom:
			vault.Denom = vault.Denom.Add(coin)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid deposit")
		}
	}

	k.SetVault(ctx, vault)
	return &types.MsgDepositResponse{
		Uid:  msg.Uid,
		Coin: msg.Coin,
	}, nil
}
