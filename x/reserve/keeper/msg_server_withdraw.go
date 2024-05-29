package keeper

import (
	"context"
	"strconv"

	"reserve/x/reserve/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Withdraw(goCtx context.Context, msg *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creator, _ := sdk.AccAddressFromBech32(msg.Creator)

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
		case vault.Status == "ready" && vault.Collateral.IsGTE(coin):
			vault.Collateral = vault.Collateral.Sub(coin)
			if vault.Collateral.Amount == sdk.ZeroInt() {
				vault.Status = "inactive"
			}
		case vault.Status == "active" && vault.Collateral.Denom == coin.Denom:
			vault.Collateral = vault.Collateral.Add(coin)
		case vault.Status == "active" && vault.Denom.Denom == coin.Denom:
			vault.Denom = vault.Denom.Add(coin)
			k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(coin))
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid withdraw")
		}
	} else {
		return nil, sdkerrors.Wrapf(types.ErrVaultNotFound, "Vault with uid %s not found", msg.Uid)
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creator, sdk.NewCoins(coin))
	if err != nil {
		return nil, err
	}

	// TODO: Handling the message
	_ = ctx

	return &types.MsgWithdrawResponse{}, nil
}
